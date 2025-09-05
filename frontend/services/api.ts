/**
 * API service for the Lilo outfit finder application
 * Handles Supabase authentication and integrates with backend client
 */
import { User } from "../types";
import { supabaseClient } from "./supabaseClient";
import { backendClient } from "./backendClient";
import { devLog, devError } from "../utils/devLog";

/**
 * Helper function to map Supabase user to our User type
 */
function mapSupabaseUser(supabaseUser: any): User {
  if (!supabaseUser) return null as any;

  return {
    id: supabaseUser.id,
    email: supabaseUser.email || "",
    name: supabaseUser.user_metadata?.name || "",
    picture: supabaseUser.user_metadata?.avatar_url || "",
    createdAt: new Date(supabaseUser.created_at),
    updatedAt: new Date(supabaseUser.updated_at || supabaseUser.created_at),
  };
}

/**
 * Helper function to check if user is authenticated
 */
async function checkAuthStatus() {
  const {
    data: { session },
    error,
  } = await supabaseClient.auth.getSession();

  devLog("Auth status check", {
    hasSession: !!session,
    sessionId: session?.user?.id,
    error,
  });

  return { session, error };
}

/**
 * API service object with methods for different endpoints
 */
const api = {
  // Auth endpoints using Supabase
  auth: {
    signup: async (email: string, password: string) => {
      devLog("Signing up user:", email);

      const { data, error } = await supabaseClient.auth.signUp({
        email,
        password,
      });

      if (error) {
        devError("Signup failed:", error);
        return {
          error: error.message,
          status: 400,
        };
      }

      // Create user profile in our backend
      if (data.user) {
        devLog("Creating user profile in backend");
        const backendResponse = await backendClient.user.createProfile({
          supabaseId: data.user.id,
          email: data.user.email!,
        });

        if (backendResponse.error) {
          devError("Backend profile creation failed:", backendResponse.error);
        }
      }

      return {
        data: {
          token: data.session?.access_token,
          user: mapSupabaseUser(data.user),
        },
        status: 201,
      };
    },

    signin: async (email: string, password: string) => {
      devLog("Signing in user:", email);

      const { data, error } = await supabaseClient.auth.signInWithPassword({
        email,
        password,
      });

      if (error) {
        devError("Signin failed:", error);
        return {
          error: error.message,
          status: 401,
        };
      }

      devLog("Signin successful");
      return {
        data: {
          token: data.session?.access_token,
          user: mapSupabaseUser(data.user),
        },
        status: 200,
      };
    },

    signout: async () => {
      devLog("Signing out user");

      const { error } = await supabaseClient.auth.signOut();

      if (error) {
        devError("Signout failed:", error);
        return {
          error: error.message,
          status: 500,
        };
      }

      devLog("Signout successful");
      return {
        status: 200,
      };
    },

    getUser: async () => {
      const { data, error } = await supabaseClient.auth.getUser();

      if (error || !data.user) {
        devError("Get user failed:", error);
        return {
          error: error?.message || "User not found",
          status: 401,
        };
      }

      // Get additional user data from our backend
      const userResponse = await backendClient.user.getProfile();

      if (userResponse.error) {
        devLog("Backend profile not found, returning Supabase user only");
        return {
          data: mapSupabaseUser(data.user),
          status: 200,
        };
      }

      return {
        data: {
          ...mapSupabaseUser(data.user),
          ...userResponse.data,
        },
        status: 200,
      };
    },

    resetPassword: async (email: string) => {
      devLog("Resetting password for:", email);

      const { error } = await supabaseClient.auth.resetPasswordForEmail(email);

      if (error) {
        devError("Password reset failed:", error);
        return {
          error: error.message,
          status: 400,
        };
      }

      return {
        message: "Password reset email sent",
        status: 200,
      };
    },
  },

  // Backend endpoints - delegate to backend client
  users: backendClient.user,
  wardrobe: backendClient.wardrobe,
  outfits: backendClient.outfits,
  recommendations: backendClient.recommendations,
  reflections: backendClient.reflections,
  wishlist: backendClient.wishlist,

  // Backend health check
  healthCheck: () => backendClient.healthCheck(),
};

export default api;
export { checkAuthStatus, backendClient };
