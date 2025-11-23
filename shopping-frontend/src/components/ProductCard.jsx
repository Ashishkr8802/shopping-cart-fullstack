export default function ProductCard({ item, onAddToCart }) {
  return (
    <div className="product-card">
      <h3>{item.name}</h3>
      <p>Status: {item.status}</p>
      <button onClick={() => onAddToCart(item.id)}>Add to Cart</button>
    </div>
  );
}
