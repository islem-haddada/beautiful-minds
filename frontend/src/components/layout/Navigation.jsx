import React from 'react';
import { Link } from 'react-router-dom';
import './Navigation.css';

const Navigation = () => {
  return (
    <nav className="navigation">
      <ul className="nav-list">
        <li><Link to="/">Accueil</Link></li>
        <li><Link to="/events">Événements</Link></li>
        <li><Link to="/about">À propos</Link></li>
        <li><Link to="/register" className="nav-cta">S'inscrire</Link></li>
        <li><Link to="/admin" className="nav-admin">🔧 Admin</Link></li>
      </ul>
    </nav>
  );
};

export default Navigation;