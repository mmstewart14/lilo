/**
 * Login screen for the Lilo outfit finder application
 */
import { useState } from "react";
import {
  Image,
  KeyboardAvoidingView,
  Platform,
  SafeAreaView,
  ScrollView,
  StyleSheet,
  TouchableOpacity,
  View,
} from "react-native";
import { LoginForm } from "../components/auth/LoginForm";
import { ThemedText } from "../components/ThemedText";
import { useThemeColor } from "../hooks/useThemeColor";

export default function LoginScreen() {
  const [isLogin, setIsLogin] = useState(true);

  const backgroundColor = useThemeColor(
    { light: "#ffffff", dark: "#121212" },
    "background"
  );

  const handleLoginSuccess = () => {
    // No need to navigate here as _layout.tsx will handle navigation
    // based on authentication state
    console.log("Login successful");
  };

  const handleRegisterPress = () => {
    setIsLogin(false);
  };

  const handleLoginPress = () => {
    setIsLogin(true);
  };

  const handleForgotPasswordPress = () => {
    // TODO: Navigate to forgot password screen (to be implemented)
    console.log("Navigate to forgot password");
  };

  return (
    <SafeAreaView style={[styles.container, { backgroundColor }]}>
      <KeyboardAvoidingView
        style={styles.keyboardAvoid}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        keyboardVerticalOffset={Platform.OS === "ios" ? 0 : 20}
      >
        <ScrollView
          contentContainerStyle={styles.scrollContent}
          showsVerticalScrollIndicator={false}
        >
          <View style={styles.logoContainer}>
            <Image
              source={require("../assets/images/icon.png")}
              style={styles.logo}
            />
            <ThemedText style={styles.appName}>Lilo</ThemedText>
            <ThemedText style={styles.tagline}>
              Your Personal Outfit Finder
            </ThemedText>
          </View>

          <View style={styles.formContainer}>
            {isLogin ? (
              <LoginForm
                onSuccess={handleLoginSuccess}
                onRegisterPress={handleRegisterPress}
                onForgotPasswordPress={handleForgotPasswordPress}
              />
            ) : (
              <View style={styles.registerContainer}>
                <ThemedText style={styles.registerTitle}>
                  Create Account
                </ThemedText>
                <ThemedText style={styles.registerSubtitle}>
                  Sign up to get started with Lilo
                </ThemedText>

                {/* Register form would go here */}
                <ThemedText style={styles.comingSoon}>
                  Registration coming soon!
                </ThemedText>

                <TouchableOpacity onPress={handleLoginPress}>
                  <ThemedText style={styles.switchFormText}>
                    Already have an account? Sign In
                  </ThemedText>
                </TouchableOpacity>
              </View>
            )}
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  keyboardAvoid: {
    flex: 1,
  },
  scrollContent: {
    flexGrow: 1,
    justifyContent: "center",
    padding: 16,
  },
  logoContainer: {
    alignItems: "center",
    marginBottom: 40,
  },
  logo: {
    width: 100,
    height: 100,
    resizeMode: "contain",
    marginBottom: 16,
  },
  appName: {
    fontSize: 32,
    fontWeight: "bold",
    marginBottom: 8,
  },
  tagline: {
    fontSize: 16,
    opacity: 0.7,
  },
  formContainer: {
    width: "100%",
  },
  registerContainer: {
    padding: 16,
  },
  registerTitle: {
    fontSize: 24,
    fontWeight: "700",
    marginBottom: 8,
  },
  registerSubtitle: {
    fontSize: 16,
    marginBottom: 24,
    opacity: 0.7,
  },
  comingSoon: {
    fontSize: 16,
    fontStyle: "italic",
    textAlign: "center",
    marginVertical: 20,
  },
  switchFormText: {
    textAlign: "center",
    marginTop: 20,
    fontSize: 14,
    fontWeight: "500",
    color: "#007aff",
  },
});
