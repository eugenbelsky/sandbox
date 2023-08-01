document.getElementById('uploadForm').addEventListener('submit', function (event) {
    event.preventDefault(); // Prevent form submission

    const imageInput = document.getElementById('imageInput');
    const imageFile = imageInput.files[0];

    if (!imageFile) {
        alert('Please select an image to upload.');
        return;
    }

    // Create a FormData object to send the image file
    const formData = new FormData();
    formData.append('image', imageFile);

    // Send the image file to the backend using fetch API
    fetch('/upload', {
        method: 'POST',
        body: formData,
    })
    .then(response => response.text())
    .then(message => {
        // Display the response from the backend to the user
        document.getElementById('responseMessage').innerText = message;
    })
    .catch(error => {
        // Handle errors, e.g., network issues or server errors
        console.error('Error uploading image:', error);
    });
});