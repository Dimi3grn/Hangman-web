function updateLives(lives) {
    const imageContainer = document.querySelector('.image-container');
    imageContainer.setAttribute('data-lives', lives);
  }
  
  // Example usage: If the player has 3 lives
  updateLives(1);
  