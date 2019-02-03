require "cannon"

function Tank(id, x, y)
  local t = {}
  t.id = id
  t.x  = x
  t.y  = y
  t.angle = 0
  t.pointOfRotation = Vec2(0, 0)  -- Offset from middle
  t.velMag = 0 -- Magnitude of acceleration
  t.w = 50
  t.h = 30
  t.cannon = Cannon(t)

  function t:getMiddle()
    return self.x+(self.w/2), self.y+(self.h/2)
  end

  function t:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      midX, midY = self:getMiddle()
      lg.translate(midX+self.pointOfRotation.x, midY+self.pointOfRotation.y)
      lg.rotate(self.angle)
      lg.translate(-self.pointOfRotation.x, -self.pointOfRotation.y)
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
      lg.line(self.h-10, 0, self.h, 0)  -- For knowing orientation.

      self.cannon:draw()
    lg.pop()
  end

  function t:applyResistance(dt)  -- Due to friction/air resistance
    self.velMag = self.velMag*(TANK_DECEL*math.min(144/dt, 1))
  end

  function t:accelerate(dt, totalPower)   -- Can be used to do both tracks rather than doing it individually.
    self.velMag = self.velMag + (TANK_ACC*dt)*totalPower
  end

  function t:leftTrack(dt, power)
    self.pointOfRotation = Vec2(self.w/2, self.h/2)
    if power > 0 then
      self.angle = self.angle + (TANK_ROTATE_RATE*dt)*math.abs(power)
    else
      self.angle = self.angle - (TANK_ROTATE_RATE*dt)*math.abs(power) -- If reversing
    end
    self:accelerate(dt, power)
    self.pointOfRotation = Vec2(0, 0)
  end

  function t:rightTrack(dt, power)
    self.pointOfRotation = Vec2(-self.w/2, -self.h/2)
    if power > 0 then
      self.angle = self.angle - (TANK_ROTATE_RATE*dt)*math.abs(power)
    else
      self.angle = self.angle + (TANK_ROTATE_RATE*dt)*math.abs(power) -- If reversing
    end
    self:accelerate(dt, power)
    self.pointOfRotation = Vec2(0, 0)
  end

  function t:update(dt)
    if love.keyboard.isDown("w", "up") then
      if love.keyboard.isDown("a", "left") then
        self:leftTrack(dt, 0.25)
        self:rightTrack(dt, 0.75)
      elseif love.keyboard.isDown("d", "right") then
        self:leftTrack(dt, 0.75)
        self:rightTrack(dt, 0.25)
      else
        self:accelerate(dt, 1)
      end
    elseif love.keyboard.isDown("s", "down") then
      if love.keyboard.isDown("a", "left") then
        self:leftTrack(dt, -1)
        self:rightTrack(dt, 1)
      elseif love.keyboard.isDown("d", "right") then
        self:leftTrack(dt, 1)
        self:rightTrack(dt, -1)
      else
        self:accelerate(dt, -1)
      end
    elseif love.keyboard.isDown("a", "left") and love.keyboard.isDown("d", "right") then
      self:accelerate(dt, 1)
    elseif love.keyboard.isDown("a", "left") then
      self:rightTrack(dt, 1)
    elseif love.keyboard.isDown("d", "right") then
      self:leftTrack(dt, 1)
    end

    self:applyResistance(dt)

    self.x = self.x + self.velMag*math.cos(self.angle)
    self.y = self.y + self.velMag*math.sin(self.angle)
    print(self.velMag)
    self.cannon:update()
  end

  table.insert(entities.players, t)
  return t
end
