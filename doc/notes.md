# Notes


## Capacity

   Battery capacity is measured in milliamps × hours (mAH). For
   example, if a battery has 250 mAH capacity and provides 2 mA
   average current to a load, in theory, the battery will last 125
   hours.

   For example, the iPhone Pro 14 battery is rated as 3,200 mAh. When
   fully charged, the battery can hold a 3,200 mA charge. If the phone
   consumes an average of 200 mA per hour, it should be able to run
   for about 16 hours. If the usage is closer to 100 mA per hour, the
   phone should be able to run for nearly 32 hours.


## Battery Condition

   • **State of Charge (SOC)(%)** – An expression of the present battery
   capacity as a percentage of maximum capacity. SOC is generally
   calculated using current integration to determine the change in
   battery capacity over time.

   • **Depth of Discharge (DOD) (%)** – The percentage of battery capacity
   that has been discharged expressed as a percentage of maximum
   capacity. A discharge to at least 80 % DOD is referred to as a deep
   discharge.

   • **Terminal Voltage (V)** – The voltage between the battery terminals
   with load applied.  Terminal voltage varies with SOC and
   discharge/charge current.

   • **Open-circuit voltage (V)** – The voltage between the battery
   terminals with no load applied. The open-circuit voltage depends on
   the battery state of charge, increasing with state of charge.


## C-Rate

   **C-rate** is a measure of the rate at which a battery is charged or
   discharged relative to its capacity. It is the charge or discharge
   current in Amps divided by the cell capacity in Ampere-hours.

   A 1C rate means that the discharge current will discharge the entire
   battery in 1 hour.

   For example, for a battery with a capacity of 500 mAh, a discharge
   rate of 5000 mA (i.e., 5 A) corresponds to a C-rate of 10C, meaning
   that such a current can discharge 10 such batteries in one
   hour. Likewise, for the same battery a charge current of 250 mA
   corresponds to a C-rate of C/2, meaning that this current will
   increase the state of charge of this battery by 50% in one hour.


## Battery Specs

   • **Nominal Voltage (V)** – The reported or reference voltage of the
   battery, also sometimes thought of as the “normal” voltage of the
   batte

   • **Capacity or Nominal Capacity (Ah for a specific C-rate)** – The
   coulometric capacity, the total Amp-hours available when the
   battery is discharged at a certain discharge current (specified as
   a C-rate) from 100 percent state-of-charge to the cut-off
   voltage. Capacity is calculated by multiplying the discharge
   current (in Amps) by the discharge time (in hours) and decreases
   with increasing C-rate.

   • **Energy or Nominal Energy (Wh (for a specific C-rate))** – The “energy
   capacity” of the battery, the total Watt-hours available when the
   battery is discharged at a certain discharge current (specified as
   a C-rate) from 100 percent state-of-charge to the cut-off
   voltage. Energy is calculated by multiplying the discharge power
   (in Watts) by the discharge time (in hours). Like capacity, energy
   decreases with increasing C-rate.

   • **Cycle Life (number for a specific DOD)** – The number of
   discharge-charge cycles the battery can experience before it fails
   to meet specific performance criteria. Cycle life is estimated for
   specific charge and discharge conditions. The actual operating life
   of the battery is affected by the rate and depth of cycles and by
   other conditions such as temperature and humidity. The higher the
   DOD, the lower the cycle life.


## Milliampere-hour vs. watt-hour

   In some cases, a battery's capacity will be given in **watt-hours**
   (Wh) rather than milliampere-hours, which is especially true for
   laptop batteries. A watt-hour is a unit of energy equal to 1 watt
   of power expended over a one-hour period. Like mAh, the Wh metric
   provides a guideline for how long a fully charged battery should
   last under normal circumstances.

   If you know a battery's voltage (V) and watt-hour rating, but not
   its mAh, you can calculate the mAh using the following formula:

    mAh = Wh x 1,000 / V

   For example, if a laptop contains a 11.4 V battery that's rated at
   75 Wh, the capacity would be about 6,579 mAh. Unfortunately, many
   laptop manufacturers list only the watt-hour rating and little
   more, in which case, you'll have to dig deeper to find the actual
   mAh.


## SysFS (Linux)


### Charge/Energy/Capacity - How to not confuse

   Because both "charge" (µAh) and "energy" (µWh) represents
   "capacity" of battery, this class distinguish these terms. Don't
   mix them!

 - **CHARGE_*** attributes represents capacity in µAh only.
 - **ENERGY_*** attributes represents capacity in µWh only.
 - **CAPACITY** attribute represents capacity in *percents*, from 0 to 100.


### Postfixes

 - **_AVG** - *hardware* averaged value, use it if your hardware is really able to report averaged values.
 - **_NOW** - momentary/instantaneous values.


### Definitions

 - **STATUS** - this attribute represents operating status (charging, full, discharging (i.e. powering a load), etc.). This corresponds to
 - **BATTERY_STATUS_*** values, as defined in battery.h.
 - **CHARGE_TYPE** - batteries can typically charge at different rates. This defines trickle and fast charges.  For batteries that are already charged or discharging, 'n/a' can be displayed (or 'unknown', if the status is not known).
 - **AUTHENTIC** - indicates the power supply (battery or charger) connected to the platform is authentic(1) or non authentic(0).
 - **HEALTH** - represents health of the battery, values corresponds to **POWER_SUPPLY_HEALTH_***, defined in battery.h.
 - **VOLTAGE_OCV** - open circuit voltage of the battery.
 - **VOLTAGE_MAX_DESIGN**, **VOLTAGE_MIN_DESIGN** - design values for maximal and minimal power supply voltages. Maximal/minimal means values of voltages when battery considered "full"/"empty" at normal conditions. Yes, there is no direct relation between voltage and battery capacity, but some dumb batteries use voltage for very approximated calculation of capacity. Battery driver also can use this attribute just to inform userspace about maximal and minimal voltage thresholds of a given battery.
 - **VOLTAGE_MAX**, **VOLTAGE_MIN** - same as **_DESIGN** voltage values except that these ones should be used if hardware could only guess (measure and retain) the thresholds of a given power supply.
 - **VOLTAGE_BOOT** - Reports the voltage measured during boot
 - **CURRENT_BOOT** - Reports the current measured during boot
 - **CHARGE_FULL_DESIGN**, **CHARGE_EMPTY_DESIGN** - design charge values, when battery considered full/empty.
 - **ENERGY_FULL_DESIGN**, **ENERGY_EMPTY_DESIGN** - same as above but for energy.
 - **CHARGE_FULL**, **CHARGE_EMPTY** - These attributes means "last remembered value of charge when battery became full/empty". It also could mean "value of charge when battery considered full/empty at given conditions (temperature, age)". I.e. these attributes represents real thresholds, not design values.
 - **ENERGY_FULL**, **ENERGY_EMPTY** - same as above but for energy.
 - **CHARGE_COUNTER** - the current charge counter (in µAh).  This could easily be negative; there is no empty or full value.  It is only useful for relative, time-based measurements.
 - **PRECHARGE_CURRENT** - the maximum charge current during precharge phase of charge cycle (typically 20% of battery capacity).
 - **CHARGE_TERM_CURRENT** - Charge termination current. The charge cycle terminates when battery voltage is above recharge threshold, and charge current is below this setting (typically 10% of battery capacity).
 - **CONSTANT_CHARGE_CURRENT** - constant charge current programmed by charger.
 - **CONSTANT_CHARGE_CURRENT_MAX** - maximum charge current supported by the power supply object.
 - **CONSTANT_CHARGE_VOLTAGE** - constant charge voltage programmed by charger.
 - **CONSTANT_CHARGE_VOLTAGE_MAX** - maximum charge voltage supported by the power supply object.
 - **INPUT_CURRENT_LIMIT** - input current limit programmed by charger. Indicates the current drawn from a charging source.
 - **CHARGE_CONTROL_LIMIT** - current charge control limit setting
 - **CHARGE_CONTROL_LIMIT_MAX** - maximum charge control limit setting
 - **CALIBRATE** - battery or coulomb counter calibration status
 - **CAPACITY** - capacity in percents.
 - **CAPACITY_ALERT_MIN** - minimum capacity alert value in percents.
 - **CAPACITY_ALERT_MAX** - maximum capacity alert value in percents.
 - **CAPACITY_LEVEL** - capacity level. This corresponds to
 - **POWER_SUPPLY_CAPACITY_LEVEL_***.
 - **TEMP_*** - temperature of the power supply.
 - **TEMP_ALERT_MIN** - minimum battery temperature alert.
 - **TEMP_ALERT_MAX** - maximum battery temperature alert.
 - **TEMP_AMBIENT_*** - ambient temperature.
 - **TEMP_AMBIENT_ALERT_MIN** - minimum ambient temperature alert.
 - **TEMP_AMBIENT_ALERT_MAX** - maximum ambient temperature alert.
 - **TEMP_MIN** - minimum operatable temperature
 - **TEMP_MAX** - maximum operatable temperature
 - **TIME_TO_EMPTY** - seconds left for battery to be considered empty (i.e. while battery powers a load)
 - **TIME_TO_FULL** - seconds left for battery to be considered full (i.e. while battery is charging)


## References

   - [Linux Kernel Documentation: Power Supply Class](https://www.kernel.org/doc/Documentation/power/power_supply_class.txt)
   - [A Guide to Understanding Battery Specifications; MIT Electric Team, December 2008](http://web.mit.edu/evt/summary_battery_specifications.pdf)

