import { useEffect, useState } from "react";
import api from "../api/client";

export default function Orders() {
  const [orders, setOrders] = useState([]);

  const fetchOrders = async () => {
    const res = await api.get("/orders");
    setOrders(res.data);
  };

  useEffect(() => {
    fetchOrders();
  }, []);

  return (
    <div className="page">
      <h2>Your Orders</h2>
      {orders.length === 0 && <p>No orders placed yet.</p>}

      <ul>
        {orders.map((order) => (
          <li key={order.id}>
            Order #{order.id} — Cart #{order.cart_id} — User #{order.user_id}
          </li>
        ))}
      </ul>
    </div>
  );
}
