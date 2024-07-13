document.addEventListener('DOMContentLoaded', function() {
    fetch('/user', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            if (data.username) {
                document.getElementById('username').textContent = data.username;
                document.getElementById('sex').textContent = data.sex;
                document.getElementById('email').textContent = data.email;
                document.getElementById('phone').textContent = data.phone;

                document.getElementById('edit-username').value = data.username;
                document.getElementById('edit-sex').value = data.sex;
                document.getElementById('edit-email').value = data.email;
                document.getElementById('edit-phone').value = data.phone;
            } else {
                alert('加载用户信息失败: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('加载用户信息时出错.');
        });
});

    function showEditForm() {
    document.getElementById('edit-form').style.display = 'block';
}

    function submitEditForm() {
    const username = document.getElementById('edit-username').value;
    const sex = document.getElementById('edit-sex').value;
    const email = document.getElementById('edit-email').value;
    const phone = document.getElementById('edit-phone').value;

    fetch('/update', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json'
},
    body: JSON.stringify({ username, sex, email, phone })
})
    .then(response => response.json())
    .then(data => {
    if (data.success) {
    alert('更新用户信息成功');
} else {
    alert('更新用户信息成功');
    document.getElementById('username').textContent = username;
    document.getElementById('sex').textContent = sex;
    document.getElementById('email').textContent = email;
    document.getElementById('phone').textContent = phone;
    document.getElementById('edit-form').style.display = 'none';
}
});
}
