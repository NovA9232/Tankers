function Cannon(parent)
  local c = {}
  c.parent = parent
  c.x = parent.x
  c.y = parent.y
  c.angle = 0
  c.w = 20
  c.h = 30

  function c:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      lg.rotate(self.angle-self.parent.angle)
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
    lg.pop()
  end

  function c:update(dt)
    self.x = self.parent.x
    self.y = self.parent.y
    mouseX, mouseY = love.mouse.getPosition()
    angleToMouse = math.atan2((mouseY - self.y), (mouseX - self.x))
    if self.angle < angleToMouse then
      self.angle = self.angle + (TANK_CANNON_ROTATE_RATE*dt)
    else
      self.angle = self.angle - (TANK_CANNON_ROTATE_RATE*dt)
    end
  end

  return c
end
