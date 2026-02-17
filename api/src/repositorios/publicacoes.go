package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {

	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ? ,?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {

	linha, erro := repositorio.db.Query(`
		SELECT 
			p.*,
			u.nick
		FROM publicacoes p inner join usuarios u
		ON u.id = p.autor_id where p.id = ?
	`, publicacaoID,
	)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {

		if erro := linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}
	return publicacao, nil

}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {

	linhas, erro := repositorio.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM publicacoes p
		INNER JOIN usuarios u
			ON u.id = p.autor_id
		LEFT JOIN seguidores s
			ON p.autor_id = s.usuario_id
		WHERE p.autor_id = ?
		OR s.seguidor_id = ?
		ORDER BY p.id DESC
`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {

		var publicacao models.Publicacao

		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorios Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {

	statement, erro := repositorios.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil

}

func (repositorios Publicacoes) Deletar(publicacaoID uint64) error {

	statement, erro := repositorios.db.Prepare("delete from publicacoes where id = ? ")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil

}

func (repositorios Publicacoes) BuscarPorUsuario(usuarioId uint64) ([]models.Publicacao, error) {

	linhas, erro := repositorios.db.Query(`
		SELECT p.*,
			u.nick
		FROM   publicacoes p
			JOIN usuarios u
				ON u.id = p.autor_id
		WHERE  p.autor_id = ? 	
	`, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {

		var publicacao models.Publicacao

		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorios Publicacoes) Curtir(publicacaoID uint64) error {

	statement, erro := repositorios.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil

}

func (repositorios Publicacoes) Descurtir(publicacaoID uint64) error {

	statement, erro := repositorios.db.Prepare(`
    UPDATE publicacoes
    SET    curtidas = CASE
                        WHEN curtidas > 0 THEN curtidas - 1
                        ELSE 0
					END
    WHERE  id = ?
    `)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
