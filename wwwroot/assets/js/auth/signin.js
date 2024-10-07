$(document).ready(function () {
});

$('#btnSignin').click(function () {
    const username = $('#username').val();
    const password = $('#password').val();

    // Buat objek data untuk dikirim
    const data = {
        username: username,
        password: password
    };

    $.ajax({
        url: '/signin',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(data),
        success: function (response) {
        },
        error: function (xhr, status, error) {
            // Menangani error
            console.error('Error:', error);

            // Mengambil pesan error dari response
            let errorMessage;
            if (xhr.responseJSON && xhr.responseJSON.errorMessage) {
                errorMessage = xhr.responseJSON.errorMessage; // Mengambil pesan error dari JSON
            } else {
                errorMessage = 'An unknown error occurred. Please try again.'; // Pesan default
            }

            alert('Error: ' + errorMessage);
        }
    });
});