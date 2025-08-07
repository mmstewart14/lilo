/**
 * OnboardingLayout component for the Lilo outfit finder application
 * Provides a consistent layout for all onboarding screens
 */
import { PropsWithChildren } from "react";
import {
  SafeAreaView,
  ScrollView,
  StyleSheet,
  TouchableOpacity,
  View,
} from "react-native";
import { ThemedText } from "../../components/ThemedText";
import { ThemedView } from "../../components/ThemedView";
import { useOnboarding } from "../../contexts/OnboardingContext";
import { useThemeColor } from "../../hooks/useThemeColor";

interface OnboardingLayoutProps extends PropsWithChildren {
  title: string;
  subtitle?: string;
  currentStep: number;
  totalSteps: number;
  onNext?: () => void;
  onBack?: () => void;
  onSkip?: () => void;
  nextDisabled?: boolean;
  nextLabel?: string;
  showSkip?: boolean;
}

export function OnboardingLayout({
  children,
  title,
  subtitle,
  currentStep,
  totalSteps,
  onNext,
  onBack,
  onSkip,
  nextDisabled = false,
  nextLabel = "Next",
  showSkip = true,
}: OnboardingLayoutProps) {
  const { isStepCompleted, resetOnboarding } = useOnboarding();
  const primaryColor = useThemeColor(
    { light: "#6200ee", dark: "#bb86fc" },
    "primary"
  );
  const backgroundColor = useThemeColor(
    { light: "#f5f5f5", dark: "#121212" },
    "background"
  );
  const secondaryColor = useThemeColor(
    { light: "#03dac6", dark: "#03dac6" },
    "secondary"
  );

  // Create step indicators
  const renderStepIndicators = () => {
    const indicators = [];
    for (let i = 1; i <= totalSteps; i++) {
      const isActive = i === currentStep;
      const isCompleted = isStepCompleted(i);

      indicators.push(
        <View
          key={i}
          style={[
            styles.stepIndicator,
            isActive && { backgroundColor: primaryColor },
            isCompleted && !isActive && { backgroundColor: secondaryColor },
            !isActive && !isCompleted && { backgroundColor: "#e0e0e0" },
          ]}
        >
          <ThemedText
            style={[
              styles.stepIndicatorText,
              (isActive || isCompleted) && { color: "white" },
            ]}
          >
            {i}
          </ThemedText>
        </View>
      );
    }
    return indicators;
  };

  return (
    <SafeAreaView style={[styles.container, { backgroundColor }]}>
      <View style={styles.header}>
        {onBack ? (
          <TouchableOpacity style={styles.backButton} onPress={onBack}>
            <ThemedText style={styles.backButtonText}>← Back</ThemedText>
          </TouchableOpacity>
        ) : (
          <View style={styles.backButton} />
        )}

        <View style={styles.progressContainer}>
          <ThemedText style={styles.progressText}>
            Step {currentStep} of {totalSteps}
          </ThemedText>
          <View style={styles.progressBarContainer}>
            <View
              style={[
                styles.progressBar,
                {
                  width: `${(currentStep / totalSteps) * 100}%`,
                  backgroundColor: primaryColor,
                },
              ]}
            />
          </View>
        </View>

        {showSkip && onSkip && (
          <TouchableOpacity style={styles.skipButton} onPress={onSkip}>
            <ThemedText style={styles.skipButtonText}>Skip</ThemedText>
          </TouchableOpacity>
        )}
      </View>

      <View style={styles.stepIndicatorsContainer}>
        {renderStepIndicators()}
      </View>

      <ScrollView
        style={styles.content}
        contentContainerStyle={styles.contentContainer}
        showsVerticalScrollIndicator={false}
      >
        <ThemedText style={styles.title}>{title}</ThemedText>
        {subtitle && (
          <ThemedText style={styles.subtitle}>{subtitle}</ThemedText>
        )}

        <ThemedView style={styles.childrenContainer}>{children}</ThemedView>
      </ScrollView>

      <View style={styles.footer}>
        <TouchableOpacity
          style={[
            styles.nextButton,
            {
              backgroundColor: nextDisabled ? "#cccccc" : primaryColor,
            },
          ]}
          onPress={onNext}
          disabled={nextDisabled}
        >
          <ThemedText style={styles.nextButtonText}>{nextLabel}</ThemedText>
        </TouchableOpacity>

        {/* TODO: For Development only */}
        <TouchableOpacity
          style={[styles.nextButton]}
          onPress={resetOnboarding}
          disabled={nextDisabled}
        >
          <ThemedText style={styles.nextButtonText}>Reset</ThemedText>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  header: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    paddingHorizontal: 16,
    paddingTop: 16,
    paddingBottom: 8,
  },
  backButton: {
    width: 60,
  },
  backButtonText: {
    fontSize: 16,
  },
  progressContainer: {
    flex: 1,
    alignItems: "center",
  },
  progressText: {
    fontSize: 14,
    marginBottom: 4,
  },
  progressBarContainer: {
    width: "100%",
    height: 4,
    backgroundColor: "#e0e0e0",
    borderRadius: 2,
    overflow: "hidden",
  },
  progressBar: {
    height: "100%",
  },
  skipButton: {
    width: 60,
    alignItems: "flex-end",
  },
  skipButtonText: {
    fontSize: 16,
  },
  stepIndicatorsContainer: {
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
    paddingVertical: 12,
  },
  stepIndicator: {
    width: 32,
    height: 32,
    borderRadius: 16,
    backgroundColor: "#e0e0e0",
    justifyContent: "center",
    alignItems: "center",
    marginHorizontal: 8,
  },
  stepIndicatorText: {
    fontSize: 14,
    fontWeight: "bold",
  },
  content: {
    flex: 1,
  },
  contentContainer: {
    padding: 24,
  },
  title: {
    fontSize: 28,
    fontWeight: "bold",
    marginBottom: 8,
  },
  subtitle: {
    fontSize: 16,
    opacity: 0.7,
    marginBottom: 24,
  },
  childrenContainer: {
    marginTop: 16,
  },
  footer: {
    padding: 24,
    paddingBottom: 32,
  },
  nextButton: {
    paddingVertical: 16,
    borderRadius: 12,
    alignItems: "center",
    justifyContent: "center",
  },
  nextButtonText: {
    color: "white",
    fontSize: 16,
    fontWeight: "600",
  },
});
