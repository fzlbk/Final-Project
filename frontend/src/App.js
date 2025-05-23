import { useState, useEffect } from "react";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState("");
  const [message, setMessage] = useState("");

  const [brandFilter, setBrandFilter] = useState("");
  const [products, setProducts] = useState([]);

  const apiUser = "http://localhost:8081";
  const apiProduct = "http://localhost:8080";

  const register = async () => {
    try {
      const res = await fetch(`${apiUser}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();
      setMessage(res.ok ? "✅ Registered successfully" : "❌ " + (data.error || "Registration failed"));
    } catch {
      setMessage("❌ Network error");
    }
  };

  const login = async () => {
    try {
      const res = await fetch(`${apiUser}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();
      if (data.token) {
        setToken(data.token);
        setMessage("✅ Logged in successfully");
      } else {
        setMessage("❌ " + (data.error || "Login failed"));
      }
    } catch {
      setMessage("❌ Network error");
    }
  };

  const fetchProducts = async () => {
    try {
      const res = await fetch(`${apiProduct}/products?brand=${brandFilter}`);
      const data = await res.json();
      setProducts(data);
    } catch {
      setMessage("❌ Failed to fetch products");
    }
  };

  useEffect(() => {
    if (token) fetchProducts();
  }, [brandFilter, token]);

  return (
    <div style={{ padding: 30, fontFamily: "Arial", maxWidth: 800 }}>
      <h2>🎵 Music Store</h2>

      {/* Auth Section */}
      <div style={{ marginBottom: 30 }}>
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          style={{ marginRight: 10, padding: 6 }}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          style={{ marginRight: 10, padding: 6 }}
        />
        <button onClick={register}>Register</button>
        <button onClick={login}>Login</button>
        <div style={{ marginTop: 10 }}>{message}</div>
      </div>

      {/* Protected Section */}
      {token ? (
        <>
          {/* Filter */}
          <div style={{ marginBottom: 20 }}>
            <input
              type="text"
              placeholder="Filter by brand"
              value={brandFilter}
              onChange={(e) => setBrandFilter(e.target.value)}
              style={{ marginRight: 10, padding: 6 }}
            />
            <button onClick={fetchProducts}>Search</button>
          </div>

          {/* Product List */}
          {products.length === 0 ? (
            <p>No products found.</p>
          ) : (
            <div style={{ display: "flex", flexWrap: "wrap", gap: 10 }}>
              {products.map((p) => (
                <div key={p.id} style={{ border: "1px solid #ccc", padding: 10, width: 250 }}>
                  <h3>{p.name}</h3>
                  <p>{p.description}</p>
                  <strong>€{p.price}</strong>
                  <div><small>Brand: {p.brand}</small></div>
                </div>
              ))}
            </div>
          )}
        </>
      ) : (
        <p style={{ color: "#888" }}>🔒 Please login to search and view products.</p>
      )}
    </div>
  );
}

export default App;
