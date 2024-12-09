import React from 'react';
import './AboutPage.css'; // Import the CSS file

const AboutPage = () => {
  const roadmap = [
    { week: 'Weeks 1-2', task: 'Initial Planning and Setup' },
    { week: 'Weeks 3-4', task: 'Develop Token Contracts' },
    { week: 'Weeks 5-6', task: 'Develop Swap Contract' },
    { week: 'Weeks 7-8', task: 'Implement Backend Listener' },
    { week: 'Weeks 9-10', task: 'Testing and Validation' },
    { week: 'Weeks 11-12', task: 'Final Deployment and Documentation' },
    { week: 'Weeks 13-14', task: 'Security Audits and Optimizations' },
    { week: 'Weeks 15-16', task: 'Advanced Features and Final Review' },
  ];

  return (
    <div className="about-container">
      <h1 className="about-title">Project Roadmap</h1>
      <div className="roadmap-container">
        {roadmap.map((item, index) => (
          <div key={index} className="roadmap-item">
            <div className="roadmap-week">{item.week}</div>
            <div className="roadmap-task">{item.task}</div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default AboutPage;
