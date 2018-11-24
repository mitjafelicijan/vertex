window.addEventListener('load', () => {
  let products = document.querySelector('ul');

  fetch('/api/products')
    .then((resp) => resp.json())
    .then(function (data) {
      data.forEach(product => {
        let item = document.createElement('li');
        item.innerText = `${product.title} (${product.price})`;
        products.appendChild(item);
      });
    });
});