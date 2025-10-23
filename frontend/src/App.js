import React, { useState, useEffect } from 'react';
import TaskList from './components/TaskList';
import AddTask from './components/AddTask';

function App() {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch tasks from API
  const fetchTasks = async () => {
    try {
      setLoading(true);
      const response = await fetch('http://localhost:8080/api/tasks');
      if (!response.ok) {
        throw new Error('Failed to fetch tasks');
      }
      const data = await response.json();
      setTasks(data);
      setError(null);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  // Add new task
  const addTask = async (title) => {
    try {
      const response = await fetch('http://localhost:8080/api/tasks', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title, completed: false }),
      });
      
      if (!response.ok) {
        throw new Error('Failed to add task');
      }
      
      const newTask = await response.json();
      setTasks([...tasks, newTask]);
    } catch (err) {
      setError(err.message);
    }
  };

  // Update task
  const updateTask = async (id, updates) => {
    try {
      console.log('Updating task:', id, updates); // Debug log
      
      const response = await fetch(`http://localhost:8080/api/tasks/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(updates),
      });
      
      console.log('Response status:', response.status); // Debug log
      
      if (!response.ok) {
        const errorText = await response.text();
        console.error('Error response:', errorText); // Debug log
        throw new Error(`Failed to update task: ${response.status}`);
      }
      
      const updatedTask = await response.json();
      console.log('Updated task:', updatedTask); // Debug log
      
      setTasks(tasks.map(task => 
        task.id === id ? updatedTask : task
      ));
    } catch (err) {
      console.error('Update task error:', err); // Debug log
      setError(err.message);
    }
  };

  // Delete task
  const deleteTask = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/api/tasks/${id}`, {
        method: 'DELETE',
      });
      
      if (!response.ok) {
        throw new Error('Failed to delete task');
      }
      
      setTasks(tasks.filter(task => task.id !== id));
    } catch (err) {
      setError(err.message);
    }
  };

  // Toggle task completion
  const toggleTask = (id) => {
    const task = tasks.find(t => t.id === id);
    if (task) {
      updateTask(id, { ...task, completed: !task.completed });
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  return (
    <div className="container">
      <div className="header">
        <h1>📝 Todo App</h1>
        <p>Full-Stack React + Go Application</p>
      </div>
      
      <div className="todo-app">
        {error && (
          <div className="error">
            Error: {error}
          </div>
        )}
        
        <AddTask onAddTask={addTask} />
        
        {loading ? (
          <div className="loading">Loading tasks...</div>
        ) : (
          <TaskList 
            tasks={tasks}
            onToggleTask={toggleTask}
            onUpdateTask={updateTask}
            onDeleteTask={deleteTask}
          />
        )}
      </div>
    </div>
  );
}

export default App;
