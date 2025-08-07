/**
 * StylePreferenceScreen for the onboarding flow
 * Allows users to select their preferred styles
 */
import { useState } from "react";
import {
  FlatList,
  Image,
  StyleSheet,
  TouchableOpacity,
  View,
} from "react-native";
import { ThemedText } from "../../components/ThemedText";
import {
  STYLE_CATEGORIES,
  useOnboarding,
} from "../../contexts/OnboardingContext";
import { useThemeColor } from "../../hooks/useThemeColor";
import { OnboardingLayout } from "./OnboardingLayout";

// Mock images for style categories
const STYLE_IMAGES: Record<string, any> = {
  Casual: require("../../assets/images/icon.png"),
  Professional: require("../../assets/images/icon.png"),
  Formal: require("../../assets/images/icon.png"),
  Athletic: require("../../assets/images/icon.png"),
  Bohemian: require("../../assets/images/icon.png"),
  Vintage: require("../../assets/images/icon.png"),
  Minimalist: require("../../assets/images/icon.png"),
  Streetwear: require("../../assets/images/icon.png"),
  Preppy: require("../../assets/images/icon.png"),
  Edgy: require("../../assets/images/icon.png"),
};

interface StylePreferenceScreenProps {
  onNext: () => void;
  onSkip: () => void;
}

export function StylePreferenceScreen({
  onNext,
  onSkip,
}: StylePreferenceScreenProps) {
  const { state, setPreferredStyles } = useOnboarding();
  const [selectedStyles, setSelectedStyles] = useState<string[]>(
    state.preferredStyles
  );

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

  const toggleStyle = (style: string) => {
    if (selectedStyles.includes(style)) {
      setSelectedStyles(selectedStyles.filter((s) => s !== style));
    } else {
      setSelectedStyles([...selectedStyles, style]);
    }
  };

  const handleNext = () => {
    setPreferredStyles(selectedStyles);
    onNext();
  };

  const renderStyleItem = ({ item }: { item: string }) => {
    const isSelected = selectedStyles.includes(item);

    return (
      <TouchableOpacity
        style={[
          styles.styleCard,
          {
            backgroundColor: cardBgColor,
            borderColor: isSelected ? primaryColor : borderColor,
          },
        ]}
        onPress={() => toggleStyle(item)}
      >
        <View
          style={[
            styles.styleImageContainer,
            { borderColor: isSelected ? primaryColor : borderColor },
          ]}
        >
          <Image source={STYLE_IMAGES[item]} style={styles.styleImage} />
        </View>
        <ThemedText style={styles.styleText}>{item}</ThemedText>
        {isSelected && (
          <View style={[styles.checkmark, { backgroundColor: primaryColor }]}>
            <ThemedText style={styles.checkmarkText}>✓</ThemedText>
          </View>
        )}
      </TouchableOpacity>
    );
  };

  return (
    <OnboardingLayout
      title="What's your style?"
      subtitle="Select all the styles that you like. This will help us recommend outfits that match your preferences."
      currentStep={1}
      totalSteps={3}
      onNext={handleNext}
      onSkip={onSkip}
      nextDisabled={selectedStyles.length === 0}
    >
      <FlatList
        data={STYLE_CATEGORIES}
        renderItem={renderStyleItem}
        keyExtractor={(item) => item}
        numColumns={2}
        columnWrapperStyle={styles.columnWrapper}
        scrollEnabled={false}
      />

      <ThemedText style={styles.helperText}>
        Select at least one style to continue
      </ThemedText>
    </OnboardingLayout>
  );
}

const styles = StyleSheet.create({
  columnWrapper: {
    justifyContent: "space-between",
  },
  styleCard: {
    width: "48%",
    borderRadius: 12,
    padding: 12,
    marginBottom: 16,
    alignItems: "center",
    borderWidth: 2,
    position: "relative",
  },
  styleImageContainer: {
    width: 80,
    height: 80,
    borderRadius: 40,
    borderWidth: 1,
    justifyContent: "center",
    alignItems: "center",
    marginBottom: 8,
    overflow: "hidden",
  },
  styleImage: {
    width: "100%",
    height: "100%",
    resizeMode: "cover",
  },
  styleText: {
    fontSize: 16,
    fontWeight: "500",
    textAlign: "center",
  },
  checkmark: {
    position: "absolute",
    top: -8,
    right: -8,
    width: 24,
    height: 24,
    borderRadius: 12,
    justifyContent: "center",
    alignItems: "center",
  },
  checkmarkText: {
    color: "white",
    fontSize: 14,
    fontWeight: "bold",
  },
  helperText: {
    fontSize: 14,
    opacity: 0.6,
    textAlign: "center",
    marginTop: 16,
  },
});
