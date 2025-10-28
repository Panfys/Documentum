import { useState } from 'react';
import styles from './AuthContainer.module.css';
import AuthInputBox from './AuthInputBox';
import { IconDoc } from './icon';
import { CircleAlert } from 'lucide-react';
import IconCheckbox from '../buttons/iconCheckBox';
import TextButton from '../buttons/textButton';
import AuthSelectBox from './AuthSelectBox';

const AuthContainer = () => {
    const [activePanel, setActivePanel] = useState('login');
    const [formData, setFormData] = useState({
        login: '',
        password: '',
        name: '',
        country: '',
        repassword: '',
        remember: false
    });

    const switchToRegister = (e) => {
        e.preventDefault();
        setActivePanel('register');
        // Очищаем форму при переключении
        setFormData({
            login: '',
            password: '',
            name: '',
            country: '',
            repassword: '',
            remember: false
        });
    };

    const switchToLogin = (e) => {
        e.preventDefault();
        setActivePanel('login');
        // Очищаем форму при переключении
        setFormData({
            login: '',
            password: '',
            name: '',
            country: '',
            repassword: '',
            remember: false
        });
    };

    const handleInputChange = (e) => {
        const { name, value, type, checked } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : value
        }));
    };

    const handleCheckboxChange = (name, checked) => {
        setFormData(prev => ({
            ...prev,
            [name]: checked
        }));
    };

    const handleSubmitLogin = (e) => {
        e.preventDefault();
        console.log('Login data:', formData);
        // Здесь будет логика входа
    };

    const handleSubmitRegister = (e) => {
        e.preventDefault();
        console.log('Register data:', formData);
        // Здесь будет логика регистрации
    };

    return (
        <div className={styles.authContainer}>
            <div className={styles.logo}>
                <IconDoc className={styles.icon} />
                <span>Documentum</span>
            </div>

            <div className={styles.authPanel}>
                {/* Панель входа */}
                <form
                    id="loginPanel"
                    className={`${styles.loginPanel} ${activePanel !== 'login' ? styles.hidden : ''}`}
                    onSubmit={handleSubmitLogin}
                >
                    <h1>Вход в систему</h1>

                    <AuthInputBox
                        label="Логин"
                        name="login"
                        type="text"
                        tabIndex={activePanel === 'login' ? 1 : -1}
                        placeholder="user"
                        value={formData.login}
                        onChange={handleInputChange}
                    />

                    <AuthInputBox
                        label="Пароль"
                        name="password"
                        type="password"
                        tabIndex={activePanel === 'login' ? 2 : -1}
                        placeholder="password"
                        value={formData.password}
                        onChange={handleInputChange}
                    />

                    <div className={styles.loginCheckButtonBox}>
                        <IconCheckbox
                            checBoxClass={styles.loginCheckbox}
                            text={"Запомнить меня"}
                            tabIndex={activePanel === 'login' ? 3 : -1}
                            name={"remember"}
                            checked={formData.remember}
                            onChange={(checked) => handleCheckboxChange('remember', checked)}
                        />

                        <TextButton
                            type="submit"
                            tabIndex={activePanel === 'login' ? 4 : -1}
                        >
                            Войти
                        </TextButton>
                    </div>

                    <div className={styles.regLink}>
                        <CircleAlert height={'1.2rem'} width={'1.2rem'} />
                        <span>Забыли пароль?</span>
                        <button
                            onClick={switchToRegister}
                            tabIndex={activePanel === 'login' ? 5 : -1}
                            type="button"
                        >
                            Зарегистрироваться
                        </button>
                    </div>
                </form>

                {/* Панель регистрации */}
                <form
                    id="registrPanel"
                    className={`${styles.loginPanel} ${activePanel !== 'register' ? styles.hidden : ''}`}
                    onSubmit={handleSubmitRegister}
                >
                    <h1>Регистрация в системе</h1>

                    <AuthInputBox
                        label="Логин"
                        name="login"
                        type="text"
                        tabIndex={activePanel === 'register' ? 1 : -1}
                        placeholder="user"
                        value={formData.login}
                        onChange={handleInputChange}
                    />

                    <AuthInputBox
                        label="Фамилия и инициалы"
                        name="name"
                        type="text"
                        tabIndex={activePanel === 'register' ? 2 : -1}
                        placeholder="Панфилов А.П."
                        value={formData.name}
                        onChange={handleInputChange}
                    />

                    <AuthSelectBox
                        label="Должность"
                        name="country"
                        options={[
                            { value: 'ru', label: 'Россия' },
                            { value: 'us', label: 'США' },
                            { value: 'de', label: 'Германия' }
                        ]}
                        value={formData.country}
                        onChange={handleInputChange}
                        tabIndex={activePanel === 'register' ? 3 : -1}
                    />

                    <AuthSelectBox
                        label="Подразделение"
                        name="country"
                        options={[
                            { value: 'ru', label: 'Россия' },
                            { value: 'us', label: 'США' },
                            { value: 'de', label: 'Германия' }
                        ]}
                        value={formData.country}
                        onChange={handleInputChange}
                        tabIndex={activePanel === 'register' ? 3 : -1}
                    />

                    <AuthSelectBox
                        label="Структурное подразделение"
                        name="country"
                        options={[
                            { value: 'ru', label: 'Россия' },
                            { value: 'us', label: 'США' },
                            { value: 'de', label: 'Германия' }
                        ]}
                        value={formData.country}
                        onChange={handleInputChange}
                        tabIndex={activePanel === 'register' ? 3 : -1}
                    />

                    <AuthInputBox
                        label="Пароль"
                        name="password"
                        type="password"
                        tabIndex={activePanel === 'register' ? 4 : -1}
                        placeholder="password"
                        value={formData.password}
                        onChange={handleInputChange}
                    />

                    <AuthInputBox
                        label="Повторите пароль"
                        name="repassword"
                        type="password"
                        tabIndex={activePanel === 'register' ? 5 : -1}
                        placeholder="password"
                        value={formData.repassword}
                        onChange={handleInputChange}
                    />

                    <TextButton
                        type="submit"
                        tabIndex={activePanel === 'register' ? 6 : -1}
                    >
                        Зарегистрироваться
                    </TextButton>

                    <div className={styles.regLink}>
                        <CircleAlert height={'1.2rem'} width={'1.2rem'} />
                        <span>Есть аккаунт?</span>
                        <button
                            onClick={switchToLogin}
                            tabIndex={activePanel === 'register' ? 7 : -1}
                            type="button"
                        >
                            Войти в систему
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default AuthContainer;