/**
 * Supabase client configuration for the Lilo outfit finder application
 */
import AsyncStorage from "@react-native-async-storage/async-storage";
import { createClient } from "@supabase/supabase-js";
import * as Constants from "expo-constants";

// Supabase configuration
let supabaseUrl = process.env.EXPO_PUBLIC_SUPABASE_URL ?? "";
let supabaseAnonKey = process.env.EXPO_PUBLIC_SUPABASE_ANON_KEY ?? "";

// Try to get values from Constants
try {
  const expoConstants = Constants.default.expoConfig;
  if (expoConstants?.extra) {
    supabaseUrl = expoConstants.extra.supabaseUrl || supabaseUrl;
    supabaseAnonKey = expoConstants.extra.supabaseAnonKey || supabaseAnonKey;
  }
} catch (error) {
  console.warn("Could not load Supabase configuration from Constants:", error);
}

// Create a single supabase client for interacting with your database
export const supabaseClient = createClient(supabaseUrl, supabaseAnonKey, {
  auth: {
    storage: AsyncStorage, // Use AsyncStorage for session persistence
    autoRefreshToken: true,
    persistSession: true,
    detectSessionInUrl: false, // Important for Expo/React Native
  },
});

// Export types
export type { Session, User } from "@supabase/supabase-js";
