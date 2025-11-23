import { useEffect, useState } from "react";
import api from "../api/client";
import ProductCard from "../components/ProductCard";
import { useAuth } from "../context/AuthContext";

export default function Products() {
  const [items, setItems] = useState([]);
  const [filtered, setFiltered] = useState([]);
  const [message, setMessage] = useState("");
  const [search, setSearch] = useState("");
  const [loading, setLoading] = useState(true);

  const { user } = useAuth();

  const fetchItems = async () => {
    try {
      setLoading(true);
      const res = await api.get("/items");
      setItems(res.data);
      setFiltered(res.data);
    } catch (err) {
      console.error(err);
      setMessage("Failed to load products. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchItems();
  }, []);

  // Search filter
  useEffect(() => {
    const q = search.toLowerCase();
    setFiltered(
      items.filter((item) => item.name.toLowerCase().includes(q))
    );
  }, [search, items]);

  const handleAddToCart = async (itemId) => {
    if (!user) {
      setMessage("Please login to add items to your cart.");
      return;
    }
    try {
      await api.post("/carts", { items: [itemId] });
      setMessage("Item added to cart ✅");
      setTimeout(() => setMessage(""), 2000);
    } catch (err) {
      console.error(err);
      setMessage("Failed to add item to cart.");
    }
  };

  return (
    <div className="page home-page">
      {/* Hero section */}
      <section className="home-hero">
        <div>
          <p className="home-badge">Shopping Portal</p>
          <h1>
            Find your <span>favourite products</span> in one place.
          </h1>
          <p className="home-subtitle">
            Browse items, add them to your cart, and checkout in just a few
            clicks. Simple, fast, and built with React + Go.
          </p>
        </div>
        <div className="home-hero-card">
          <p className="hero-label">Current user</p>
          {user ? (
            <>
              <p className="hero-username">{user.username}</p>
              <p className="hero-text">You’re logged in and ready to shop 🛒</p>
            </>
          ) : (
            <p className="hero-text">
              You’re browsing as a guest. <br />
              <span>Login or Signup to start adding items to your cart.</span>
            </p>
          )}
        </div>
      </section>

      {/* Search / actions */}
      <div className="home-toolbar">
        <div className="home-search">
          <span className="search-icon">🔍</span>
          <input
            type="text"
            placeholder="Search products by name..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
        </div>
        <button className="ghost-btn" onClick={fetchItems}>
          ⟳ Refresh
        </button>
      </div>

      {message && <p className="info">{message}</p>}

      {/* Products grid */}
      {loading ? (
        <div className="products-grid">
          <div className="product-skeleton" />
          <div className="product-skeleton" />
          <div className="product-skeleton" />
        </div>
      ) : filtered.length === 0 ? (
        <div className="empty-state">
          <p className="empty-title">No products found</p>
          <p className="empty-subtitle">
            Try clearing the search or click <strong>Refresh</strong>.
          </p>
        </div>
      ) : (
        <div className="products-grid">
          {filtered.map((item) => (
            <ProductCard
              key={item.id}
              item={item}
              onAddToCart={handleAddToCart}
            />
          ))}
        </div>
      )}
    </div>
  );
}
