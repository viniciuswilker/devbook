$('#nova-publicacao').on('submit', criarPublicacao)

function criarPublicacao(e) {
    e.preventDefault()

    console.log("titulo:", $('#titulo').val())
    console.log("conteudo:", $('#conteudo').val())


    $.ajax({
        url: "/publicacoes",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        })
    }).done(function () {
        window.location = "/home"
    }).fail(function () {
        alert("erro ao criar")
    })
}




$(document).on('click', '.curtir-publicacao', curtirPublicacao)
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao)
$(document).on('click', '#atualizar-publicacao', atualizarPublicacao)


function curtirPublicacao(e) {
    e.preventDefault()

    const elemento = $(this)
    const publicacaoId = elemento
        .closest('[data-publicacao-id]')
        .data('publicacao-id')

    elemento.css('pointer-events', 'none')

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST",
    }).done(function () {

        const contador = elemento.find('.contador-curtidas')
        const quantidade = parseInt(contador.text())

        contador.text(quantidade + 1)

        elemento.removeClass('curtir-publicacao')
        elemento.addClass('descurtir-publicacao text-danger')

    }).fail(function () {
        alert('erro ao curtir')
    }).always(function () {
        elemento.css('pointer-events', 'auto')
    })
}

function descurtirPublicacao(e) {
    e.preventDefault()

    const elemento = $(this)
    const publicacaoId = elemento
        .closest('[data-publicacao-id]')
        .data('publicacao-id')

    elemento.css('pointer-events', 'none')

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST",
    }).done(function () {

        const contador = elemento.find('.contador-curtidas')
        const quantidade = parseInt(contador.text())

        contador.text(Math.max(0, quantidade - 1))

        elemento.removeClass('descurtir-publicacao text-danger')
        elemento.addClass('curtir-publicacao')

    }).fail(function () {
        alert('erro ao descurtir')
    }).always(function () {
        elemento.css('pointer-events', 'auto')
    })
}


function atualizarPublicacao(e) {
    $(this).prop('disabled', true)


    const publicacaoId = $(this).data('publicacao-id')

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        contentType: "application/json",
        data: JSON.stringify({
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        })
    }).done(function () {
        alert('Publicação atualizada com sucesso')
    }).fail(function () {
        alert('Erro ao atualizar')
    }).always(function () {
        $('#atualizar-publicacao').prop('disabled', false)
    })

}