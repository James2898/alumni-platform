import { useState } from 'react';
import { Navbar } from './components/Navbar';
import { DirectoryPage } from './pages/DirectoryPage';
import { LoginPage } from './pages/LoginPage';
import { RegisterPage } from './pages/RegisterPage';
import { ProfilePage } from './pages/ProfilePage';
import { AdminPage } from './pages/AdminPage';
import { useAuth } from './context/AuthContext';

export function App() {
  const [currentTab, setCurrentTab] = useState('directory');
  const { user } = useAuth();

  return (
    <div className="min-h-screen flex flex-col bg-slate-50 text-slate-900">
      <Navbar currentTab={currentTab} onSelectTab={setCurrentTab} />

      <main className="flex-1">
        {currentTab === 'directory' && <DirectoryPage />}
        {currentTab === 'login' && (
          <LoginPage
            onSuccess={() => setCurrentTab('directory')}
            onGoToRegister={() => setCurrentTab('register')}
          />
        )}
        {currentTab === 'register' && (
          <RegisterPage
            onSuccess={() => setCurrentTab('login')}
            onGoToLogin={() => setCurrentTab('login')}
          />
        )}
        {currentTab === 'profile' && <ProfilePage />}
        {currentTab === 'admin' && (
          user?.role === 'admin' ? (
            <AdminPage />
          ) : (
            <div className="max-w-md mx-auto my-16 text-center text-slate-500">
              Access restricted to administrators.
            </div>
          )
        )}
      </main>

      <footer className="bg-white border-t border-slate-200 py-6 text-center text-xs text-slate-500">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex flex-col sm:flex-row items-center justify-between gap-4">
          <p>© 2026 Alumni Platform. Modern Go & React Cloud-Native Architecture.</p>
          <div className="flex items-center gap-4 text-slate-400">
            <span>MySQL 8.0</span>
            <span>•</span>
            <span>MinIO S3</span>
            <span>•</span>
            <span>Docker Compose</span>
            <span>•</span>
            <span>CKAD Ready</span>
          </div>
        </div>
      </footer>
    </div>
  );
}

export default App;
