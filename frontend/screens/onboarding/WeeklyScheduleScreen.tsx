/**
 * WeeklyScheduleScreen for the onboarding flow
 * Allows users to set their weekly schedule for outfit recommendations
 */
import { ScrollView, StyleSheet, TouchableOpacity, View } from "react-native";
import { ThemedText } from "../../components/ThemedText";
import { DAY_TYPES, useOnboarding } from "../../contexts/OnboardingContext";
import { useThemeColor } from "../../hooks/useThemeColor";
import { WeeklySchedule } from "../../types";
import { OnboardingLayout } from "./OnboardingLayout";

interface WeeklyScheduleScreenProps {
  onNext: () => void;
  onBack: () => void;
  onSkip: () => void;
}

export function WeeklyScheduleScreen({
  onNext,
  onBack,
  onSkip,
}: WeeklyScheduleScreenProps) {
  const { state, setWeeklySchedule } = useOnboarding();
  const { weeklySchedule } = state;

  const primaryColor = useThemeColor(
    { light: "#6200ee", dark: "#bb86fc" },
    "primary"
  );
  const cardBgColor = useThemeColor(
    { light: "#ffffff", dark: "#1e1e1e" },
    "card"
  );
  const borderColor = useThemeColor(
    { light: "#e0e0e0", dark: "#333333" },
    "border"
  );

  const days: (keyof WeeklySchedule)[] = [
    "monday",
    "tuesday",
    "wednesday",
    "thursday",
    "friday",
    "saturday",
    "sunday",
  ];

  const formatDayName = (day: string): string => {
    return day.charAt(0).toUpperCase() + day.slice(1);
  };

  const formatDayType = (type: string): string => {
    return type.charAt(0).toUpperCase() + type.slice(1);
  };

  const handleNext = () => {
    onNext();
  };

  return (
    <OnboardingLayout
      title="Weekly Schedule"
      subtitle="Tell us about your typical week so we can recommend appropriate outfits for each day."
      currentStep={2}
      totalSteps={3}
      onNext={handleNext}
      onBack={onBack}
      onSkip={onSkip}
    >
      <ScrollView style={styles.container}>
        {days.map((day) => (
          <View key={day} style={styles.dayContainer}>
            <ThemedText style={styles.dayText}>{formatDayName(day)}</ThemedText>
            <View style={styles.typeContainer}>
              {DAY_TYPES.map((type) => (
                <TouchableOpacity
                  key={type}
                  style={[
                    styles.typeButton,
                    {
                      backgroundColor:
                        weeklySchedule[day] === type
                          ? primaryColor
                          : cardBgColor,
                      borderColor:
                        weeklySchedule[day] === type
                          ? primaryColor
                          : borderColor,
                    },
                  ]}
                  onPress={() => setWeeklySchedule(day, type)}
                >
                  <ThemedText
                    style={[
                      styles.typeText,
                      {
                        color:
                          weeklySchedule[day] === type ? "white" : undefined,
                      },
                    ]}
                  >
                    {formatDayType(type)}
                  </ThemedText>
                </TouchableOpacity>
              ))}
            </View>
          </View>
        ))}

        <ThemedText style={styles.helperText}>
          Tap on a day type to set your schedule for each day of the week.
        </ThemedText>
      </ScrollView>
    </OnboardingLayout>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  dayContainer: {
    marginBottom: 16,
  },
  dayText: {
    fontSize: 18,
    fontWeight: "600",
    marginBottom: 8,
  },
  typeContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    marginLeft: -4,
    marginRight: -4,
  },
  typeButton: {
    paddingHorizontal: 12,
    paddingVertical: 8,
    borderRadius: 20,
    marginHorizontal: 4,
    marginBottom: 8,
    borderWidth: 1,
  },
  typeText: {
    fontSize: 14,
    fontWeight: "500",
  },
  helperText: {
    fontSize: 14,
    opacity: 0.6,
    textAlign: "center",
    marginTop: 24,
    marginBottom: 16,
  },
});
