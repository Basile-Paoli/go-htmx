import React, {useState} from "react";
import {createRoot} from "react-dom/client";

function App() {
    const [count, setCount] = useState(0);
    return (
        <div>
            <p>Test: {count}</p>
            <button onClick={() => setCount(count + 1)}>Increment</button>
        </div>
    );
}


createRoot(document.getElementById("count")!).render(<App/>);
