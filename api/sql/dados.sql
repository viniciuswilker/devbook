INSERT INTO usuarios (nome, nick, email, senha)
values
("usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$NXIf1n0Q8iBoAm2TjiLTkO4oETSDBERTu64T8uhdskL3mmwrRK2si"),
("usuario 2", "caio", "caio@email.com", "$2a$10$NXIf1n0Q8iBoAm2TjiLTkO4oETSDBERTu64T8uhdskL3mmwrRK2si"),
("usuario 3", "vinicius", "vinicius@email.com", "$2a$10$NXIf1n0Q8iBoAm2TjiLTkO4oETSDBERTu64T8uhdskL3mmwrRK2si"),
("usuario 4", "luana", "luana@email.com", "$2a$10$NXIf1n0Q8iBoAm2TjiLTkO4oETSDBERTu64T8uhdskL3mmwrRK2si"),
("usuario 5", "joao", "joao@email.com", "$2a$10$NXIf1n0Q8iBoAm2TjiLTkO4oETSDBERTu64T8uhdskL3mmwrRK2si");


insert into seguidores (usuario_id, seguidor_id)
values
(1,2),
(1,3),
(2,3),
(4,2);


insert into publicacoes (titulo, conteudo, autor_id) values
("Publi do user 1", "Essa é a publi do usuario 1!", 1),
("Publi do user 2", "Essa é a publi do usuario 2!", 2),
("Publi do user 3", "Essa é a publi do usuario 3!", 3),
("Publi do user 4", "Essa é a publi do usuario 4!", 4),
("Publi do user 2", "Essa é a publi do usuario 2!", 2);