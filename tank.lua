require "cannon"

function Tank(id, x, y)
  local t = {}
  t.id = id
  t.x  = x
  t.y  = y
  t.angle = 0
  t.pointOfRotation = Vec2(0, 0)  -- Offset from middle
  t.velMag = 0 -- Magnitude of acceleration
  t.w = 30
  t.h = 50
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

      self.cannon:draw()
    lg.pop()
  end

  function t:applyResistance(dt)  -- Due to friction/air resistance
    self.velMag = self.velMag*(TANK_DECEL*math.min(1/dt, 1))
  end

  function t:update(dt)
    if love.keyboard.isDown("w", "up") then
      self.velMag = self.velMag + (TANK_ACC*METER)*dt
    end

    if love.keyboard.isDown("a", "left") then
      t.pointOfRotation = Vec2(self.w/2, self.h/2)
      self.angle = self.angle + (TANK_ROTATE_RATE*dt)
      self.velMag = self.velMag + ((TANK_ACC*METER)*dt)/2  -- /2 since only 1 track
      t.pointOfRotation = Vec2(0, 0)
    end

    if not love.keyboard.isDown("w", "a", "s", "d", "up", "left", "down", "right") then
      self:applyResistance(dt)
    end


    self.x = self.x + self.velMag*math.cos(self.angle)
    self.y = self.y + self.velMag*math.sin(self.angle)
    --print(self.velMag*math.cos(self.angle), self.velMag*math.sin(self.angle))
    self.cannon:update()
  end

  table.insert(entities.players, t)
  return t
end
