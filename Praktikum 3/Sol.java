import java.util.Collection;
import java.util.List;
import java.util.function.Predicate;

public class Sol {

    public static <T> long countPredicate(Collection<T> collection, Predicate<T> predicate) {
        return collection.stream()
                         .filter(predicate)
                         .count();
    }

    public static <T> void swap(T[] array, int idx1, int idx2) {
        T temp = array[idx1];
        array[idx1] = array[idx2];
        array[idx2] = temp;
    }

    public static <T extends Comparable<T>> T findMax(List<T> list, int begin, int end) { {
            return list.stream()
                       .skip(begin)
                       .limit(end - begin)
                       .max(Comparable::compareTo)
                       .orElse(null);
        }
    }
}