import { getPost } from '../pages/api/data'
let pageTitle = document.title
const response = await getPost(pageTitle)
const data = await response.json()

if (!data.error) {
	document.querySelector('.edit').addEventListener('click', () => {
		// 存储文件名
		localStorage.setItem('md-title', pageTitle)
		window.location.href = `/edit`
        // 存储文件内容
        localStorage.setItem('md-content', data.content)
	})
}

document.querySelector('.delete').addEventListener('click', () => {
	const title = pageTitle	
	const form = document.querySelector('.delete-form')
   	form.querySelector('input[name="title"]').value = title
	form.submit()
})