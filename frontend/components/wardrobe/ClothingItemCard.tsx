/**
 * ClothingItemCard component for displaying clothing items in the wardrobe
 */
import { Image } from "expo-image";
import React from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { useThemeColor } from "../../hooks/useThemeColor";
import { ClothingItem } from "../../types";
import { ThemedText } from "../ThemedText";

interface ClothingItemCardProps {
  item: ClothingItem;
  onPress?: () => void;
  onEdit?: () => void;
  onDelete?: () => void;
  compact?: boolean;
}

export function ClothingItemCard({
  item,
  onPress,
  onEdit,
  onDelete,
  compact = false,
}: ClothingItemCardProps) {
  const backgroundColor = useThemeColor(
    { light: "#fff", dark: "#1c1c1e" },
    "background"
  );
  const textColor = useThemeColor({ light: "#000", dark: "#fff" }, "text");
  const borderColor = useThemeColor(
    { light: "#e0e0e0", dark: "#2c2c2e" },
    "border"
  );

  const imageUrl =
    item.imageUrls && item.imageUrls.length > 0 ? item.imageUrls[0] : null;

  return (
    <TouchableOpacity
      style={[
        styles.container,
        { backgroundColor, borderColor },
        compact ? styles.compactContainer : null,
      ]}
      onPress={onPress}
      activeOpacity={0.8}
    >
      <View
        style={[
          styles.imageContainer,
          compact ? styles.compactImageContainer : null,
        ]}
      >
        {imageUrl ? (
          <Image
            source={{ uri: imageUrl }}
            style={styles.image}
            contentFit="cover"
            transition={200}
          />
        ) : (
          <View
            style={[styles.placeholderImage, { backgroundColor: borderColor }]}
          >
            <ThemedText style={styles.placeholderText}>No Image</ThemedText>
          </View>
        )}
        {!item.isOwned && (
          <View style={styles.wishlistTag}>
            <ThemedText style={styles.wishlistText}>Wishlist</ThemedText>
          </View>
        )}
      </View>

      <View
        style={[
          styles.contentContainer,
          compact ? styles.compactContentContainer : null,
        ]}
      >
        <ThemedText style={styles.title} numberOfLines={1}>
          {item.name}
        </ThemedText>

        {!compact && (
          <>
            <View style={styles.detailsRow}>
              <ThemedText style={styles.detailLabel}>Category:</ThemedText>
              <ThemedText style={styles.detailValue}>
                {item.category}
              </ThemedText>
            </View>

            <View style={styles.detailsRow}>
              <ThemedText style={styles.detailLabel}>Color:</ThemedText>
              <ThemedText style={styles.detailValue}>{item.color}</ThemedText>
            </View>

            {item.brand && (
              <View style={styles.detailsRow}>
                <ThemedText style={styles.detailLabel}>Brand:</ThemedText>
                <ThemedText style={styles.detailValue}>{item.brand}</ThemedText>
              </View>
            )}

            <View style={styles.tagsContainer}>
              {item.season.map((season, index) => (
                <View
                  key={index}
                  style={[styles.tag, { backgroundColor: borderColor }]}
                >
                  <Text style={[styles.tagText, { color: textColor }]}>
                    {season}
                  </Text>
                </View>
              ))}
            </View>

            <View style={styles.actionsContainer}>
              {onEdit && (
                <TouchableOpacity
                  style={[styles.actionButton, { backgroundColor: "#2196F3" }]}
                  onPress={onEdit}
                >
                  <Text style={styles.actionButtonText}>Edit</Text>
                </TouchableOpacity>
              )}

              {onDelete && (
                <TouchableOpacity
                  style={[styles.actionButton, { backgroundColor: "#F44336" }]}
                  onPress={onDelete}
                >
                  <Text style={styles.actionButtonText}>Delete</Text>
                </TouchableOpacity>
              )}
            </View>
          </>
        )}
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    borderRadius: 12,
    overflow: "hidden",
    marginVertical: 8,
    borderWidth: 1,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 2,
  },
  compactContainer: {
    flexDirection: "row",
    height: 80,
  },
  imageContainer: {
    position: "relative",
    height: 180,
  },
  compactImageContainer: {
    height: 80,
    width: 80,
  },
  image: {
    width: "100%",
    height: "100%",
  },
  placeholderImage: {
    width: "100%",
    height: "100%",
    justifyContent: "center",
    alignItems: "center",
  },
  placeholderText: {
    fontSize: 14,
    fontWeight: "500",
  },
  wishlistTag: {
    position: "absolute",
    top: 10,
    right: 10,
    backgroundColor: "rgba(33, 150, 243, 0.8)",
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 4,
  },
  wishlistText: {
    color: "white",
    fontSize: 12,
    fontWeight: "600",
  },
  contentContainer: {
    padding: 16,
  },
  compactContentContainer: {
    flex: 1,
    padding: 8,
    justifyContent: "center",
  },
  title: {
    fontSize: 16,
    fontWeight: "600",
    marginBottom: 8,
  },
  detailsRow: {
    flexDirection: "row",
    marginBottom: 4,
  },
  detailLabel: {
    fontSize: 14,
    fontWeight: "500",
    marginRight: 4,
    opacity: 0.7,
  },
  detailValue: {
    fontSize: 14,
  },
  tagsContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    marginTop: 8,
    marginBottom: 12,
  },
  tag: {
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 4,
    marginRight: 8,
    marginBottom: 8,
  },
  tagText: {
    fontSize: 12,
    fontWeight: "500",
  },
  actionsContainer: {
    flexDirection: "row",
    justifyContent: "flex-end",
  },
  actionButton: {
    paddingVertical: 6,
    paddingHorizontal: 12,
    borderRadius: 4,
    marginLeft: 8,
  },
  actionButtonText: {
    color: "white",
    fontWeight: "500",
    fontSize: 12,
  },
});
