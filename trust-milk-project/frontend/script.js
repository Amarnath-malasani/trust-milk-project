document.getElementById("loginForm").onsubmit = function(event) {
    event.preventDefault();
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
  
    fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username: username, password: password }),
      credentials: "include"
    })
    .then(response => response.json())
    .then(data => {
      if (data.message) {
        alert(data.message);
        localStorage.setItem("loggedInUser", username);
        checkLoginStatus();
      } else {
        alert(data.error);
      }
    });
  };
  
  document.getElementById("logoutBtn").onclick = function() {
    fetch("http://localhost:8080/logout", { credentials: "include" })
    .then(response => response.json())
    .then(() => {
      alert("Logged out successfully!");
      localStorage.removeItem("loggedInUser");
      checkLoginStatus();
    });
  };
  
  function checkLoginStatus() {
    fetch("http://localhost:8080/api/dashboard", { credentials: "include" })
    .then(response => response.json())
    .then(data => {
      if (data.user) {
        document.getElementById("loginBtn").style.display = "none";
        document.getElementById("logoutBtn").style.display = "block";
      } else {
        document.getElementById("loginBtn").style.display = "block";
        document.getElementById("logoutBtn").style.display = "none";
      }
    });
  }
  
  window.onload = checkLoginStatus;
  