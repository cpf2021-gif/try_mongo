import { getPost } from '../pages/api/data'
let pageTitle = document.title

let currentUrl = window.location.href
// 获取最后一个/后面的内容
let id = currentUrl.substring(currentUrl.lastIndexOf('/') + 1)

const response = await getPost(id)
const data = await response.json()

if (!data.error) {
	document.querySelector('.edit').addEventListener('click', () => {
		// 存储文件名
		localStorage.setItem('md-title', pageTitle)
		window.location.href = `/edit`
        // 存储文件内容
        localStorage.setItem('md-content', data.content)
		// 存储文件id
		localStorage.setItem('md-id', id)
	})
}

document.querySelector('.delete').addEventListener('click', () => {
	const form = document.querySelector('.delete-form')
   	form.querySelector('input[name="id"]').value = id
	form.submit()
})