const div = document.getElementById("paragraphs")

fetch(`/goblog/blog`,{
    method:"GET",
    headers:{
        "Content-type":"application/json"
    }
}).then(res=>res.json())
.then(posts=>{
    /* posts.array.forEach(element => {
        console.log(element)
        const newContent = document.createTextNode(element)
        div.appendChild(newContent)
    }); */
   var  i=0
  var oldDiv=div
   console.log("post lenght",posts.length)
    for ( i = 0; i<posts.length;i++){
        console.log(posts[i].title)
        var newDiv = document.createElement(`div`)
        newDiv.className = `box`
        div.appendChild(newDiv)
        var newTitle = document.createElement("h4")
        newTitle.innerHTML= `${posts[i].title}`
        var newContent = document.createElement("p")
        newContent.innerHTML= `${posts[i].content}`
        var newCategory = document.createElement("p")
        newCategory.innerHTML=`Category: ${posts[i].category}`
        var newTags = document.createElement("p")
        newTags.innerHTML=`Tags: ${posts[i].tags}`
        var newCreated = document.createElement("p")
        newCreated.innerHTML=`Created at: ${posts[i].created_at}`
        var newUpdated = document.createElement("p")
        newUpdated.innerHTML=`Updated at: ${posts[i].updated_at}`
        newDiv.appendChild(newTitle)
        newDiv.appendChild(newContent)
        newDiv.appendChild(newCategory)
        newDiv.appendChild(newTags)
        newDiv.appendChild(newCreated)
        newDiv.appendChild(newUpdated)
        oldDiv=newDiv
    }
    console.log(JSON.stringify(posts))    
    console.log(posts)
})
