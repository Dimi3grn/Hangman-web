function switchForm(formType) {
    // Hide both forms
    document.getElementById('login-form').style.display = 'none';
    document.getElementById('signup-form').style.display = 'none';
  
    // Remove 'active' class from both buttons
    const buttons = document.querySelectorAll('.tab-button');
    buttons.forEach(button => button.classList.remove('active'));
  
    // Show the selected form
    if (formType === 'login') {
      document.getElementById('login-form').style.display = 'block';
      document.querySelector('button[onclick="switchForm(\'login\')"]').classList.add('active');
    } else if (formType === 'signup') {
      document.getElementById('signup-form').style.display = 'block';
      document.querySelector('button[onclick="switchForm(\'signup\')"]').classList.add('active');
    }
  }
  
  // Initially display the login form
  switchForm('login');
  