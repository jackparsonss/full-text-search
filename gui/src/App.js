import axios from "axios";
import { useEffect, useState } from "react";
import "./App.css";

const debounce = function (fn, t) {
  let timeOut;
  return function (...args) {
    clearTimeout(timeOut);
    timeOut = setTimeout(() => {
      fn(...args);
    }, t);
  };
};

function App() {
  const [result, setResult] = useState([]);
  const [query, setQuery] = useState("");

  const debouncedResult = debounce(setResult, 500);

  const search = async () => {
    try {
      const res = await axios.post("http://localhost:8080/search", { query });
      debouncedResult(res.data.result);
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    search();
  }, [query]);

  return (
    <div className="main">
      <div class="input-container">
        <input type="text" onChange={(e) => setQuery(e.target.value)} />
        <label>enter your query</label>
      </div>

      <div className="list-container">
        <ul>
          {result.map((document) => (
            <li>{document}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
