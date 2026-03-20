package kata;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

class LegacyCheckoutCalculatorTest {

    private final LegacyCheckoutCalculator calculator = new LegacyCheckoutCalculator();

    @ParameterizedTest(name = "{0}")
    @MethodSource("sharedCases")
    void calculatesExpectedTotals(
            String name,
            String customerType,
            int subtotalCents,
            String country,
            String couponCode,
            boolean blackFriday,
            int expectedTotalCents
    ) {
        Order order = new Order(customerType, subtotalCents, country, couponCode, blackFriday);

        assertEquals(expectedTotalCents, calculator.calculateTotalCents(order), name);
    }

    @Test
    void treatsNullStringsAsEmptyValues() {
        Order order = new Order(null, 10_000, null, null, false);

        assertEquals(12_500, calculator.calculateTotalCents(order));
    }

    private static Stream<Arguments> sharedCases() {
        return Stream.of(
                caseData("regular_it_baseline", "regular", 10_000, "IT", "", false, 12_900),
                caseData("new_de_baseline", "new", 10_000, "DE", "", false, 12_800),
                caseData("unknown_customer_other_country", "guest", 10_000, "FR", "", false, 12_500),
                caseData("premium_below_10000", "premium", 9_999, "US", "", false, 11_663),
                caseData("premium_at_10000", "premium", 10_000, "DE", "", false, 11_610),
                caseData("save10_at_threshold", "regular", 5_000, "DE", "SAVE10", false, 6_255),
                caseData("save10_below_threshold", "regular", 4_999, "DE", "SAVE10", false, 6_848),
                caseData("viponly_true_with_trim", " vip ", 18_000, " IT ", " VIPONLY ", false, 17_980),
                caseData("viponly_false_for_non_vip", "premium", 10_000, "DE", "VIPONLY", false, 11_610),
                caseData("bulk_at_threshold", "regular", 20_000, "FR", "BULK", false, 21_100),
                caseData("bulk_below_threshold", "regular", 19_999, "FR", "BULK", false, 22_499),
                caseData("black_friday_regular_us", "regular", 10_000, "US", "", true, 11_965),
                caseData("black_friday_employee_us", "employee", 10_000, "US", "", true, 9_790),
                caseData("black_friday_non_us_no_shipping_surcharge", "regular", 10_000, "DE", "", true, 12_205),
                caseData("freeship_at_threshold", "regular", 8_000, "DE", "FREESHIP", false, 9_520),
                caseData("freeship_below_threshold", "regular", 7_999, "DE", "FREESHIP", false, 10_418),
                caseData("freeship_overrides_black_friday_us_surcharge", "regular", 9_000, "US", "FREESHIP", true, 9_148),
                caseData("vip_free_shipping_exact_threshold", "vip", 17_648, "DE", "", false, 17_850),
                caseData("premium_free_shipping_exact_threshold", "premium", 22_223, "DE", "", false, 23_800),
                caseData("premium_free_shipping_just_below", "premium", 22_222, "DE", "", false, 24_698),
                caseData("taxfree_non_it", "regular", 10_000, "DE", "TAXFREE", false, 10_900),
                caseData("taxfree_it_has_no_effect", "regular", 10_000, "IT", "TAXFREE", false, 12_900),
                caseData("vip_it_tax_override", "vip", 10_000, "IT", "", false, 10_900),
                caseData("tax_us_standard_rate", "regular", 10_000, "US", "", false, 12_200),
                caseData("employee_it_no_shipping_surcharge", "employee", 10_000, "IT", "", false, 9_240),
                caseData("employee_save10_hits_40_percent_boundary", "employee", 5_000, "DE", "SAVE10", false, 4_970),
                caseData("employee_freeship_still_pays_surcharge", "employee", 12_000, "DE", "FREESHIP", false, 10_496),
                caseData("negative_total_clamped_to_zero", "regular", -10_000, "FR", "", false, 0)
        );
    }

    private static Arguments caseData(
            String name,
            String customerType,
            int subtotalCents,
            String country,
            String couponCode,
            boolean blackFriday,
            int expectedTotalCents
    ) {
        return Arguments.of(name, customerType, subtotalCents, country, couponCode, blackFriday, expectedTotalCents);
    }
}