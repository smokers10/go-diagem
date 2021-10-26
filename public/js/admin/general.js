$(document).ready(function(){
    const logged = JSON.parse(localStorage.getItem("logged"))
    let nama = logged.nama
    let role = logged.roles
    
    $("#logged-admin-name").text(nama)
    $("#role-badge").text(role)
})