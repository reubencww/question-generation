import { format, formatDistanceToNow, parseISO } from "date-fns";

function ucFirst(string: string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}

export const fuzzyDate = (date: string | null): string => {
    if (!date) {
        return "";
    }

    let formatted = formatDistanceToNow(parseISO(date), {
        addSuffix: true,
    });

    // capitalize first letter
    return ucFirst(formatted);
};

export const formatDate = (
    date: string,
    formatToFollow: string = "yyyy-MM-dd HH:mm:ss"
) => {
    return format(parseISO(date), formatToFollow);
};
