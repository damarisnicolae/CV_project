function deleteUser (userId) {
  fetch(`/user/${userId}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok')
      }
      return response.json()
    })
    .then((data) => {
      if (data.success) {
        alert('User deleted successfully')
        window.location.reload()
      } else {
        throw new Error('User deletion failed')
      }
    })
    .catch((error) => {
      alert('Error deleting user')
      console.error(
        'There was a problem with the fetch operation:',
        error.message
      )
    })
}
document.getElementById('delete-button').addEventListener('click', function () {
  const userId = this.getAttribute('data-user-id')
  deleteUser(userId)
})
