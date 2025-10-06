title = document.getElementById("title")
content= document.getElementById("content")
category=document.getElementById("category")
tags=document.getElementById("tags")
submit=document.getElementById("submit")
div = document.getElementById("box")
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
    })
    .then(response =>{
        console.log(response)
        var output = document.createElement("p")
        output.innerHTML = 'Posted!'
         var ref = document.createElement("href")
        ref.innerHTML = "<a href='http://195.251.68.21:8080/'>Return to main menu</a>"
        div.appendChild(output)
        div.appendChild(ref)
        output.innerHTML = ''
        div.appendChild(output)

       
    }).catch(error =>{
        console.log(response)
        var output = document.createElement("p")
        output.innerHTML = `${error}`
         var ref = document.createElement("href")
        ref.innerHTML = "<a href='http://195.251.68.21:8080/'>Return to main menu</a>"
        div.appendChild(output)
        div.appendChild(ref)
    })
})