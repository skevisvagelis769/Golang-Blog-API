title = document.getElementById("title")
content= document.getElementById("content")
category=document.getElementById("category")
tags=document.getElementById("tags")
submit=document.getElementById("submit")

submit.addEventListener('click',e =>{
    console.log(title.value,content.value,category.value,tags.value)
    console.log("clicked")
    fetch('http://195.251.68.21:8080/blog',{
        method: "POST",
        headers: {"Content-type":"application/json"},
        body: JSON.stringify({
            "title":`${title.value}`,
            "content":`${content.value}`,
            "category":`${category.value}`,
            "tags":`${tags.value}`
        })
    }).then(res => res.json())
    .then(response =>{
        console.log(response)
    })
})