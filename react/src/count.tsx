import React, {useState} from "react";
import {createRoot} from "react-dom/client";

function App() {
    const initialCount = (window as any).countProps?.initialCount ?? 0;
    const [count, setCount] = useState(initialCount);
    return (
        <div>
            <p>Count: {count}</p>
            <button onClick={() => setCount(count + 1)}>Increment</button>
        </div>
    );
}


createRoot(document.getElementById("count")!).render(<App/>);
