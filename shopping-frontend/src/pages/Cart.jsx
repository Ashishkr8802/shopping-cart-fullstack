import { useEffect, useState } from "react";
import api from "../api/client";

export default function Cart() {
  const [carts, setCarts] = useState([]);
  const [message, setMessage] = useState("");

  const fetchCarts = async () => {
    const res = await api.get("/carts");
    setCarts(res.data);
  };

  useEffect(() => {
    fetchCarts();
  }, []);

  const handleCheckout = async (cartId) => {
    try {
      await api.post("/orders", { cart_id: cartId });
      setMessage("Order placed successfully!");
      fetchCarts(); // refresh carts (status changes to ordered)
    } catch (err) {
      console.error(err);
      setMessage("Failed to place order.");
    }
  };

  return (
    <div className="page">
      <h2>Your Cart</h2>
      {message && <p className="info">{message}</p>}

      {carts.length === 0 && <p>No carts found.</p>}

      {carts.map((cart) => (
        <div key={cart.id} className="cart-card">
          <h3>Cart #{cart.id}</h3>
          <p>Status: {cart.status}</p>
          <ul>
            {cart.items?.map((ci) => (
              <li key={ci.id}>
                {ci.item?.name || `Item #${ci.item_id}`}
              </li>
            ))}
          </ul>
          {cart.status === "open" && (
            <button onClick={() => handleCheckout(cart.id)}>Checkout</button>
          )}
        </div>
      ))}
    </div>
  );
}
