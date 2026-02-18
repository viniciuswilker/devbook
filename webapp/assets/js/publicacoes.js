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




$(".curtir-publicacao").on('click', curtirPublicacao)
function curtirPublicacao(e) {

    e.preventDefault()

    const elementoClicado = $(e.target)
    const publicacaoId = elementoClicado.closest('[data-publicacao-id]').data('publicacao-id')

    elementoClicado.prop('disabled', true)

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST",
    }).done(function () {

        const contador = elementoClicado.find('.contador-curtidas')
        const quantidade = parseInt(contador.text())

        contador.text(quantidade + 1)

    }).fail(function () {
        alert('erro ao curtir')
    }).always(function () {
        elementoClicado.prop('disabled', false)

    })

}