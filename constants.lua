METER = 10  -- Size of 1 meter in pixels

-- Tank:
TANK_ACC = 0.1*METER  -- Tank acceleration in m/s-2
TANK_DECEL = 0.99   -- Resistive force due to friction (multiplier of current velocity.)
TANK_ROTATE_RATE = math.pi/4  -- Rotational speed of tank in radians per second.
