(() => {
  if (!window.EventSource) {
    alert("No EventSource!");
    return;
  }
  let chatlog = document.getElementById("chat-log");
  let chatmsg = document.getElementById("chat-msg");

  let isBlank = function (string) {
    return string == null || string.trim() === "";
  };
  let username;
  while (isBlank(username)) {
    username = prompt("What's your name?");
    if (!isBlank(username)) {
      document.getElementById("user-name").innerHTML = `<b>${username}</b>`;
    }
  }
  let inputform = document.getElementById("input-form");

  inputform.addEventListener("submit", (e) => {
    e.preventDefault();
    data = { msg: chatmsg.value, name: username };
    fetch("/messages", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
    },
      body: JSON.stringify(data),
    }).then((res) => {
      console.log("Request complete! response:", res);
    });
  });
  let addMessage = (data) => {
    let text = "";
    if (!isBlank(data.name)) {
      text = `<strong>${data.name}: </strong>`;
    }
    text += data.msg;
    chatlog.innerHTML+=(`<div><span>${text}</span></div>`)
  };
  addMessage({
    msg: "hello",
    name: "aaa",
  });

let es = new EventSource("/stream")
es.onopen=function(e){
    data = {name: username}
    fetch("/users", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      }).then((res) => {
        console.log("Request complete! response:", res);
      });
}
es.onmessage=(e)=>{
   let msg =  JSON.parse(e.data)
   addMessage(msg)
}
//창이 닫히기 직전에 호출
window.onbeforeunload=()=>{
    fetch(`/users?username=${username}`, {
        method: 'DELETE',
      }).then((res) => {
        es.close()
      });
}


})();
