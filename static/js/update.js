document.addEventListener('DOMContentLoaded', function() {
    fetch('/user', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            if (data.message !== "User not found") {
                document.getElementById('username').value = data.username;
                document.getElementById('sex').value = data.sex;
                document.getElementById('email').value = data.email;
                document.getElementById('phone').value = data.phone;
            } else {
                alert('Failed to load user information: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred while loading user information.');
        });
});

    document.getElementById('updateForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(event.target);

    fetch('/update', {
    method: 'POST',
    body: JSON.stringify(Object.fromEntries(formData)),
    headers: {
    'Content-Type': 'application/json'
}
})
    .then(response => response.json())
    .then(data => {
    if (data.message === 'User updated successfully') {
    alert('User information updated successfully!');
} else {
    alert('Failed to update user information: ' + data.message);
}
})
    .catch(error => {
    console.error('Error:', error);
    alert('An error occurred while updating user information.');
});
});
