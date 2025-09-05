/**
 * Development logging utility
 * Only outputs console logs in development environment
 */

/**
 * Check if we're in development mode
 */
const isDevelopment = (): boolean => {
  // For Expo/React Native
  if (typeof __DEV__ !== 'undefined') {
    return __DEV__;
  }
  
  // For Node.js environments
  if (typeof process !== 'undefined' && process.env) {
    return process.env.NODE_ENV === 'development';
  }
  
  // Fallback - assume development if we can't determine
  return true;
};

/**
 * Development console.log - only logs in development
 */
export const devLog = (...args: any[]): void => {
  if (isDevelopment()) {
    console.log(...args);
  }
};

/**
 * Development console.warn - only logs in development
 */
export const devWarn = (...args: any[]): void => {
  if (isDevelopment()) {
    console.warn(...args);
  }
};

/**
 * Development console.error - only logs in development
 */
export const devError = (...args: any[]): void => {
  if (isDevelopment()) {
    console.error(...args);
  }
};

/**
 * Export the development check function
 */
export { isDevelopment };

/**
 * Default export for convenience
 */
export default devLog;