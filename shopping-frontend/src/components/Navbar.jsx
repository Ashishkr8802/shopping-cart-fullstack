import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

export default function Navbar() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <nav className="navbar">
      <div className="nav-left">
        <Link to="/" className="logo">Shopping App</Link>
      </div>
      <div className="nav-right">
        <Link to="/">Products</Link>
        {user && (
          <>
            <Link to="/cart">Cart</Link>
            <Link to="/orders">Orders</Link>
          </>
        )}
        {!user && <Link to="/login">Login</Link>}
        {!user && <Link to="/signup">Signup</Link>}
        {user && (
          <button onClick={handleLogout} className="logout-btn">
            Logout ({user.username})
          </button>
        )}
      </div>
    </nav>
  );
}
