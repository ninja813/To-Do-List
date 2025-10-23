import React, { useState } from 'react';
import TaskItem from './TaskItem';

const TaskList = ({ tasks, onToggleTask, onUpdateTask, onDeleteTask }) => {
  const [editingId, setEditingId] = useState(null);

  const handleEdit = (id) => {
    setEditingId(id);
  };

  const handleSave = (id, newTitle) => {
    onUpdateTask(id, { title: newTitle });
    setEditingId(null);
  };

  const handleCancel = () => {
    setEditingId(null);
  };

  if (tasks.length === 0) {
    return (
      <div className="empty-state">
        <h3>🎉 No tasks yet!</h3>
        <p>Add your first task above to get started.</p>
      </div>
    );
  }

  return (
    <div className="task-list">
      {tasks.map(task => (
        <TaskItem
          key={task.id}
          task={task}
          isEditing={editingId === task.id}
          onToggle={onToggleTask}
          onEdit={handleEdit}
          onSave={handleSave}
          onCancel={handleCancel}
          onDelete={onDeleteTask}
        />
      ))}
    </div>
  );
};

export default TaskList;
