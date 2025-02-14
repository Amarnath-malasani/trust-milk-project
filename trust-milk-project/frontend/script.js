// Get the modals
var loginModal = document.getElementById("loginPopup");
var createAccountModal = document.getElementById("createAccountPopup");

// Get the buttons that open the modals
var loginBtn = document.getElementById("loginBtn");
var createAccountLink = document.getElementById("createAccountLink");

// Get the <span> elements that close the modals
var loginClose = loginModal.getElementsByClassName("close")[0];
var createAccountClose = createAccountModal.getElementsByClassName("close")[0];

// When the user clicks the login button, open the login modal
loginBtn.onclick = function() {
  loginModal.style.display = "block";
}

// When the user clicks the create account link, open the create account modal
createAccountLink.onclick = function() {
  loginModal.style.display = "none";
  createAccountModal.style.display = "block";
}

// When the user clicks on <span> (x), close the modals
loginClose.onclick = function() {
  loginModal.style.display = "none";
}

createAccountClose.onclick = function() {
  createAccountModal.style.display = "none";
}

// When the user clicks anywhere outside of the modals, close them
window.onclick = function(event) {
  if (event.target == loginModal) {
    loginModal.style.display = "none";
  }
  if (event.target == createAccountModal) {
    createAccountModal.style.display = "none";
  }
}

// Handle form submissions
document.getElementById("loginForm").onsubmit = function(event) {
  event.preventDefault();
  var username = document.getElementById("username").value;
  var password = document.getElementById("password").value;

  fetch("/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ username: username, password: password })
  })
  .then(response => response.json())
  .then(data => {
    if (data.message) {
      alert(data.message);
      loginModal.style.display = "none";
      localStorage.setItem("username", data.username);
      updateUIWithUsername(data.username);
    } else {
      alert(data.error);
    }
  });
};

document.getElementById("createAccountForm").onsubmit = function(event) {
  event.preventDefault();
  var username = document.getElementById("newUsername").value;
  var mobile = document.getElementById("mobileNumber").value;
  var password = document.getElementById("newPassword").value;
  var confirmPassword = document.getElementById("confirmPassword").value;

  if (password !== confirmPassword) {
    alert("Passwords do not match");
    return;
  }

  if (!/^\d{10}$/.test(mobile)) {
    alert("Mobile number must be exactly 10 digits");
    return;
  }

  fetch("/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ username: username, mobile: mobile, password: password })
  })
  .then(response => response.json())
  .then(data => {
    if (data.message) {
      alert(data.message);
      createAccountModal.style.display = "none";
    } else {
      alert(data.error);
    }
  });
};

function updateUIWithUsername(username) {
  var usernameDisplay = document.getElementById("usernameDisplay");
  usernameDisplay.textContent = username;
}

// Check if the user is already logged in
document.addEventListener("DOMContentLoaded", function() {
  var username = localStorage.getItem("username");
  if (username) {
    updateUIWithUsername(username);
  }
});

// Toggle password visibility
document.querySelectorAll('.toggle-password').forEach(item => {
  item.addEventListener('click', function() {
    var input = this.previousElementSibling;
    if (input.type === "password") {
      input.type = "text";
      this.classList.add('fa-eye-slash');
      this.classList.remove('fa-eye');
    } else {
      input.type = "password";
      this.classList.add('fa-eye');
      this.classList.remove('fa-eye-slash');
    }
  });
});