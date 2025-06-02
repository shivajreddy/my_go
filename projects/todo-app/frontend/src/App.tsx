// frontend/src/App.tsx
import { useEffect, useState } from 'react';

function App() {
    const [todos, setTodos] = useState<any[]>([]);
    const [title, setTitle] = useState('');

    useEffect(() => {
        fetch('/api/todos').then(res => res.json()).then(setTodos);
    }, []);

    const addTodo = () => {
        fetch('/api/todo', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title, done: false })
        }).then(() => window.location.reload());
    };

    return (
        <div>
            <h1>Todos</h1>
            <input value={title} onChange={e => setTitle(e.target.value)} />
            <button onClick={addTodo}>Add</button>
            <ul>
                {todos.map(todo => <li key={todo.id}>{todo.title}</li>)}
            </ul>
        </div>
    );
}
export default App;

