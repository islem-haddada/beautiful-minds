const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

// Fonction générique pour les requêtes GET
export const get = async (endpoint) => {
  try {
    const response = await fetch(`${API_BASE_URL}${endpoint}`);
    if (!response.ok) {
      let text = await response.text();
      try {
        const body = JSON.parse(text || '{}');
        text = body.message || JSON.stringify(body) || text;
      } catch (e) {
        // not JSON, keep text
      }
      throw new Error(`Erreur ${response.status}: ${text}`);
    }
    return await response.json();
  } catch (error) {
    console.error('Erreur GET:', error);
    throw error;
  }
};

// Fonction générique pour les requêtes POST
export const post = async (endpoint, data) => {
  try {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      let text = await response.text();
      try {
        const body = JSON.parse(text || '{}');
        text = body.message || JSON.stringify(body) || text;
      } catch (e) {
        // not JSON
      }
      throw new Error(`Erreur ${response.status}: ${text}`);
    }
    return await response.json();
  } catch (error) {
    console.error('Erreur POST:', error);
    throw error;
  }
};

// API Membres
export const memberAPI = {
  getAll: () => get('/members'),
  getById: (id) => get(`/members/${id}`),
  create: (data) => post('/members', data),
};

// API Événements
export const eventAPI = {
  getAll: () => get('/events'),
  getById: (id) => get(`/events/${id}`),
  create: (data) => post('/events', data),
  register: (eventId, memberId) => post(`/events/${eventId}/register`, { member_id: memberId }),
};

// API Annonces
export const announcementAPI = {
  getAll: () => get('/announcements'),
  getById: (id) => get(`/announcements/${id}`),
  create: (data) => post('/announcements', data),
};