/**
 * Onboarding flow for the Lilo outfit finder application
 */
import { AppRoutes } from "@/constants/AppRoutes";
import { useRouter } from "expo-router";
import { useEffect, useState } from "react";
import { ActivityIndicator, StyleSheet, View } from "react-native";
import { SafeAreaProvider } from "react-native-safe-area-context";
import { ThemedText } from "../components/ThemedText";
import {
  OnboardingProvider,
  useOnboarding,
} from "../contexts/OnboardingContext";
import { StylePreferenceScreen } from "../screens/onboarding/StylePreferenceScreen";
import { WardrobeQuickAddScreen } from "../screens/onboarding/WardrobeQuickAddScreen";
import { WeeklyScheduleScreen } from "../screens/onboarding/WeeklyScheduleScreen";

// Wrapper component that uses the onboarding context
function OnboardingFlow() {
  const router = useRouter();
  const {
    state,
    nextStep,
    prevStep,
    markStepCompleted,
    submitOnboarding,
    saveOnboardingProgress,
    isOnboardingComplete,
  } = useOnboarding();
  const { currentStep } = state;
  const [loading, setLoading] = useState(false);
  const [checkingCompletion, setCheckingCompletion] = useState(true);

  // Check if onboarding is already complete
  useEffect(() => {
    const checkOnboardingStatus = async () => {
      try {
        const completed = await isOnboardingComplete();
        if (completed) {
          router.replace(AppRoutes.Home);
        }
      } catch (error) {
        console.error("Error checking onboarding status:", error);
      } finally {
        setCheckingCompletion(false);
      }
    };

    checkOnboardingStatus();
  }, [isOnboardingComplete, router]);

  const handleNext = async () => {
    // Mark current step as completed
    markStepCompleted(currentStep);

    if (currentStep < 3) {
      nextStep();
      await saveOnboardingProgress();
    } else {
      // Onboarding complete, submit data and navigate to home
      setLoading(true);
      const success = await submitOnboarding();
      if (success) {
        router.replace("/(tabs)");
      } else {
        setLoading(false);
        // Show error message or retry option
      }
    }
  };

  const handleBack = () => {
    if (currentStep > 1) {
      prevStep();
    }
  };

  const handleSkip = async () => {
    // Skip to the next step but don't mark as completed
    if (currentStep < 3) {
      nextStep();
      await saveOnboardingProgress();
    } else {
      // Onboarding complete, submit data and navigate to home
      setLoading(true);
      const success = await submitOnboarding();
      if (success) {
        router.replace(AppRoutes.Home);
      } else {
        setLoading(false);
        // Show error message or retry option
      }
    }
  };

  const handleComplete = async () => {
    // Mark final step as completed
    markStepCompleted(currentStep);

    // Submit onboarding data
    setLoading(true);
    const success = await submitOnboarding();
    if (success) {
      router.replace(AppRoutes.Home);
    } else {
      setLoading(false);
      // Show error message or retry option
    }
  };

  if (loading || checkingCompletion) {
    return (
      <View style={styles.loadingContainer}>
        <ActivityIndicator size="large" color="#6200ee" />
        <ThemedText style={styles.loadingText}>Loading...</ThemedText>
      </View>
    );
  }

  return (
    <View style={styles.container}>
      {currentStep === 1 && (
        <StylePreferenceScreen onNext={handleNext} onSkip={handleSkip} />
      )}
      {currentStep === 2 && (
        <WeeklyScheduleScreen
          onNext={handleNext}
          onBack={handleBack}
          onSkip={handleSkip}
        />
      )}
      {currentStep === 3 && (
        <WardrobeQuickAddScreen
          onNext={handleComplete}
          onBack={handleBack}
          onSkip={handleComplete}
        />
      )}
    </View>
  );
}

export default function OnboardingScreen() {
  return (
    <SafeAreaProvider>
      <OnboardingProvider>
        <OnboardingFlow />
      </OnboardingProvider>
    </SafeAreaProvider>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  loadingContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  loadingText: {
    marginTop: 16,
    fontSize: 16,
  },
});
