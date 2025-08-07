/**
 * OnboardingRevisitModal component
 * Allows users to revisit onboarding steps after initial onboarding is complete
 */
import { useRouter } from "expo-router";
import {
  Modal,
  ScrollView,
  StyleSheet,
  TouchableOpacity,
  View,
} from "react-native";
import { useOnboarding } from "../../contexts/OnboardingContext";
import { useThemeColor } from "../../hooks/useThemeColor";
import { ThemedText } from "../ThemedText";
import { ThemedView } from "../ThemedView";

interface OnboardingRevisitModalProps {
  visible: boolean;
  onClose: () => void;
}

export function OnboardingRevisitModal({
  visible,
  onClose,
}: OnboardingRevisitModalProps) {
  const router = useRouter();
  const { goToStep, isStepCompleted } = useOnboarding();

  const primaryColor = useThemeColor(
    { light: "#6200ee", dark: "#bb86fc" },
    "primary"
  );
  const backgroundColor = useThemeColor(
    { light: "#ffffff", dark: "#121212" },
    "background"
  );
  const cardBgColor = useThemeColor(
    { light: "#f5f5f5", dark: "#1e1e1e" },
    "card"
  );
  const secondaryColor = useThemeColor(
    { light: "#03dac6", dark: "#03dac6" },
    "secondary"
  );

  const onboardingSteps = [
    {
      id: 1,
      title: "Style Preferences",
      description: "Update your style preferences and fashion choices.",
    },
    {
      id: 2,
      title: "Weekly Schedule",
      description: "Adjust your weekly schedule for outfit recommendations.",
    },
    {
      id: 3,
      title: "Wardrobe Quick Add",
      description: "Add more items to your wardrobe collection.",
    },
  ];

  const handleStepSelect = (stepId: number) => {
    goToStep(stepId);
    router.push("/onboarding");
    onClose();
  };

  return (
    <Modal
      animationType="slide"
      transparent={true}
      visible={visible}
      onRequestClose={onClose}
    >
      <View style={styles.centeredView}>
        <ThemedView style={[styles.modalView, { backgroundColor }]}>
          <View style={styles.header}>
            <ThemedText style={styles.modalTitle}>
              Revisit Onboarding
            </ThemedText>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <ThemedText style={styles.closeButtonText}>✕</ThemedText>
            </TouchableOpacity>
          </View>

          <ThemedText style={styles.modalSubtitle}>
            Select a step to revisit and update your preferences
          </ThemedText>

          <ScrollView style={styles.stepsContainer}>
            {onboardingSteps.map((step) => {
              const completed = isStepCompleted(step.id);

              return (
                <TouchableOpacity
                  key={step.id}
                  style={[
                    styles.stepCard,
                    { backgroundColor: cardBgColor },
                    completed && styles.completedStepCard,
                  ]}
                  onPress={() => handleStepSelect(step.id)}
                >
                  <View style={styles.stepHeader}>
                    <View
                      style={[
                        styles.stepIndicator,
                        {
                          backgroundColor: completed
                            ? secondaryColor
                            : "#cccccc",
                        },
                      ]}
                    >
                      <ThemedText style={styles.stepIndicatorText}>
                        {step.id}
                      </ThemedText>
                    </View>
                    <ThemedText style={styles.stepTitle}>
                      {step.title}
                    </ThemedText>
                  </View>
                  <ThemedText style={styles.stepDescription}>
                    {step.description}
                  </ThemedText>
                  {completed && (
                    <View style={styles.completedBadge}>
                      <ThemedText style={styles.completedText}>
                        Completed
                      </ThemedText>
                    </View>
                  )}
                </TouchableOpacity>
              );
            })}
          </ScrollView>

          <TouchableOpacity
            style={[styles.closeModalButton, { backgroundColor: primaryColor }]}
            onPress={onClose}
          >
            <ThemedText style={styles.closeModalButtonText}>Close</ThemedText>
          </TouchableOpacity>
        </ThemedView>
      </View>
    </Modal>
  );
}

const styles = StyleSheet.create({
  centeredView: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "rgba(0, 0, 0, 0.5)",
  },
  modalView: {
    width: "90%",
    maxHeight: "80%",
    borderRadius: 20,
    padding: 24,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
  },
  header: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    marginBottom: 16,
  },
  modalTitle: {
    fontSize: 24,
    fontWeight: "bold",
  },
  closeButton: {
    width: 32,
    height: 32,
    borderRadius: 16,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "rgba(150, 150, 150, 0.1)",
  },
  closeButtonText: {
    fontSize: 16,
    fontWeight: "bold",
  },
  modalSubtitle: {
    fontSize: 16,
    opacity: 0.7,
    marginBottom: 24,
  },
  stepsContainer: {
    marginBottom: 24,
  },
  stepCard: {
    padding: 16,
    borderRadius: 12,
    marginBottom: 16,
  },
  completedStepCard: {
    borderLeftWidth: 4,
    borderLeftColor: "#03dac6",
  },
  stepHeader: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: 8,
  },
  stepIndicator: {
    width: 28,
    height: 28,
    borderRadius: 14,
    justifyContent: "center",
    alignItems: "center",
    marginRight: 12,
  },
  stepIndicatorText: {
    color: "white",
    fontSize: 14,
    fontWeight: "bold",
  },
  stepTitle: {
    fontSize: 18,
    fontWeight: "600",
  },
  stepDescription: {
    fontSize: 14,
    opacity: 0.7,
    marginLeft: 40,
  },
  completedBadge: {
    position: "absolute",
    top: 16,
    right: 16,
    paddingHorizontal: 8,
    paddingVertical: 4,
    backgroundColor: "#03dac6",
    borderRadius: 12,
  },
  completedText: {
    color: "white",
    fontSize: 12,
    fontWeight: "500",
  },
  closeModalButton: {
    paddingVertical: 14,
    borderRadius: 12,
    alignItems: "center",
    justifyContent: "center",
  },
  closeModalButtonText: {
    color: "white",
    fontSize: 16,
    fontWeight: "600",
  },
});
