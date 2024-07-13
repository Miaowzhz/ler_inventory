// 杨士涵：注册
document.getElementById('registerForm').onsubmit = async function(event) {
    event.preventDefault();
    const form = event.target;
    const formData = new FormData(form);

    const response = await fetch(form.action, {
    method: form.method,
    body: formData
});

    if (!response.ok) {
    const data = await response.json();
    showModal(data.message);
} else {
    window.location.href = '/login';
}
};

    function showModal(message) {
    const modal = document.getElementById("myModal");
    const modalMessage = document.getElementById("modalMessage");
    const closeButtons = document.querySelectorAll(".close");

    modalMessage.innerText = message;
    modal.style.display = "block";

    closeButtons.forEach(button => {
    button.onclick = function() {
    modal.style.display = "none";
};
});

    window.onclick = function(event) {
    if (event.target == modal) {
    modal.style.display = "none";
}
};
}
