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