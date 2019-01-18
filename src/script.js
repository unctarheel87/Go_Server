(function() {

  const app = document.getElementById('app');
  const h2 = document.createElement('h2');

  fetch('http://localhost:8080/api')
    .then(res => res.json())
    .then(data => {
      const { Name, Age } = data;
      app.appendChild(h2).textContent = `My name is ${Name} and I'm ${Age} years old.`
    });

})()