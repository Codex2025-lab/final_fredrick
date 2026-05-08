const productsContainer = document.getElementById("products");
const loader = document.getElementById("loader");

let page = 1;
let limit = 20;
let loading = false;
let hasMore = true;

// Fetch products from Go backend
async function fetchProducts() {
  if (loading || !hasMore) return;

  loading = true;

  try {
    const response = await fetch(
      `http://localhost:8080/products?page=${page}&limit=${limit}`
    );

    const result = await response.json();

    displayProducts(result.data);

    // Stop if no more data
    if (result.data.length < limit) {
      hasMore = false;
      loader.innerText = "No more products";
    }

    page++;
  } catch (error) {
    console.error("Error fetching products:", error);
  }

  loading = false;
}

// Render products
function displayProducts(products) {
  products.forEach(product => {
    const div = document.createElement("div");

    div.className = "product";

    div.innerHTML = `
      <h3>${product.name}</h3>
      <p>ID: ${product.id}</p>
    `;

    productsContainer.appendChild(div);
  });
}

// Observer Pattern using IntersectionObserver
const observer = new IntersectionObserver((entries) => {
  const entry = entries[0];

  if (entry.isIntersecting) {
    fetchProducts();
  }
}, {
  threshold: 1.0
});

// Observe loader element
observer.observe(loader);

// Initial fetch
fetchProducts();