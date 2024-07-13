// 当页面加载完毕时，获取收藏的待办事项
document.addEventListener('DOMContentLoaded', function() {
    fetchFavorites();
});

// 孙文乐：获取收藏的待办事项
function fetchFavorites() {
    fetch('/v1/todo?favorite=true')
        .then(response => response.json())
        .then(data => {
            const tbody = document.getElementById('favorites-table-body');
            tbody.innerHTML = '';
            data.forEach((todo, index) => {
                if (todo.favorite) {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                            <td>${index + 1}</td>
                            <td>${todo.title}</td>
                            <td><button onclick="toggleFavorite(${todo.id})">取消收藏</button></td>
                        `;
                    tbody.appendChild(row);
                }
            });
        })
        .catch(error => console.error('获取收藏的待办事项时出错:', error));
}

// 切换收藏状态
function toggleFavorite(id) {
    fetch(`/v1/todo/favorite/${id}`, { method: 'PUT' })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                fetchFavorites();  // 切换后刷新列表
            } else {
                console.error('切换收藏状态时出错:', data.error);
            }
        })
        .catch(error => console.error('切换收藏状态时出错:', error));
}