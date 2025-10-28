import { useState, useRef, useEffect } from 'react';

export function useTogglePanel(initialState = false) {
    const [isExpanded, setIsExpanded] = useState(initialState);
    const panelRef = useRef(null);
    const buttonRef = useRef(null);

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (panelRef.current && 
                !panelRef.current.contains(event.target) &&
                !(buttonRef.current && buttonRef.current.contains(event.target))) {
                setIsExpanded(false);
            }
        };

        document.addEventListener('mousedown', handleClickOutside);
        return () => document.removeEventListener('mousedown', handleClickOutside);
    }, []);

    const togglePanel = () => setIsExpanded(!isExpanded);

    return {
        isExpanded,
        togglePanel,
        panelRef,
        buttonRef
    };
}